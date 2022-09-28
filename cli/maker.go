package main

import (
	"bytes"
	"embed"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var (
	timeLength  = 10
	placeholder = strings.Repeat("-", timeLength)
	//go:embed templates/*
	templateFS embed.FS
	repeats    = make(map[string]struct{})
)

type Maker struct {
	tp   *template.Template
	pg   *Program
	libp string
	out  bool
	err  error
}

func NewMaker(pg *Program) *Maker {
	path, _ := exec.LookPath("go")
	m := &Maker{pg: pg, libp: filepath.Dir(path) + string(os.PathSeparator)}
	m.tp, m.err = template.ParseFS(
		templateFS,
		"templates/cmd/*",
		"templates/internal/biz/*",
		"templates/internal/conf/*",
		"templates/internal/data/*",
		"templates/internal/service/*",
	)
	return m
}

func (m *Maker) RemoveAll(path string) {
	if m.err != nil {
		return
	}

	before := time.Now()
	m.pg.Output("\x1b[1;33m[\x1b[1;5;33m%s\x1b[0m\x1b[1;33m]\x1b[0m remove \x1b[1;4;34m[%s]\x1b[0m", placeholder, path)
	m.out = true
	m.err = os.RemoveAll(path)
	if m.err == nil {
		m.pg.Output("\x1b[1A\x1b[1;33m[%10s]\x1b[0m remove \x1b[1;34m[%s]\x1b[0m", m.dur(time.Since(before)), path)
	}
}

func (m *Maker) parse(
	path, key string,
	kind ast.ObjKind,
	vf func([]byte, *bytes.Buffer, *ast.CallExpr),
	ff func([]byte, *bytes.Buffer, *ast.FuncDecl),
) {
	if m.err != nil {
		return
	}

	var src []byte
	src, m.err = os.ReadFile(path)
	if m.err != nil {
		return
	}

	fset := token.NewFileSet()
	var f *ast.File
	f, m.err = parser.ParseFile(fset, "", src, parser.DeclarationErrors)
	if m.err != nil {
		return
	}
	for k, obj := range f.Scope.Objects {
		if m.err != nil {
			return
		}

		if k != key || obj.Kind != kind {
			continue
		}

		before := time.Now()
		buf := &bytes.Buffer{}
		switch d := obj.Decl.(type) {
		case *ast.ValueSpec:
			for _, v := range d.Values {
				ce, ok := v.(*ast.CallExpr)
				if !ok {
					continue
				}

				se, ok := ce.Fun.(*ast.SelectorExpr)
				if !ok {
					continue
				}

				if se.X.(*ast.Ident).Name != "wire" && se.Sel.Name != "NewSet" {
					continue
				}

				vf(src, buf, ce)
			}
		case *ast.FuncDecl:
			ff(src, buf, d)
		}

		if buf.Len() > 0 {
			m.err = os.WriteFile(path, buf.Bytes(), 0744)
		} else {
			m.pg.Output("\x1b[1A\x1b[1;33m[%10s]\x1b[0m can't find service with file \x1b[1;34m[%s]\x1b[0m, will be skip.",
				m.dur(time.Since(before)), path)
		}
	}
}

func (m *Maker) RevertService(path, key, name string, kind ast.ObjKind) {
	if m.err != nil {
		return
	}

	before := time.Now()
	if _, ok := repeats[path]; !ok {
		m.pg.Output("\x1b[1;33m[\x1b[1;5;33m%s\x1b[0m\x1b[1;33m]\x1b[0m remove service from \x1b[1;4;34m[%s]\x1b[0m", placeholder, path)
		m.out = true
	}

	m.parse(path, key, kind, func(src []byte, buf *bytes.Buffer, expr *ast.CallExpr) {
		if len(expr.Args) == 0 {
			return
		}

		var idx int
		var r token.Pos
		for i, arg := range expr.Args {
			service := arg.(*ast.Ident)
			if service.Name == name {
				idx = i
				r = arg.End()
				break
			}
		}

		if idx == 0 {
			buf.Write(src[:expr.Lparen])
			buf.Write(src[r:])
			return
		}

		buf.Write(src[:expr.Args[idx-1].End()])
		buf.Write(src[r:])
	}, func(src []byte, buf *bytes.Buffer, d *ast.FuncDecl) {
		if len(d.Type.Params.List) == 0 {
			return
		}

		var idx int
		for i, p := range d.Type.Params.List {
			if p.Type.(*ast.StarExpr).X.(*ast.Ident).Name == name {
				idx = i
				break
			}
		}

		writeParams := func(idx int, fs []*ast.Field) {
			for _, p := range fs {
				buf.WriteString("\n\t")
				buf.WriteString("s" + strconv.Itoa(idx) + " *")
				buf.WriteString(p.Type.(*ast.StarExpr).X.(*ast.Ident).Name)
				buf.WriteString(",")
				idx++
			}
		}

		writeBody := func(idx, count int) {
			for i := 0; i < count; i++ {
				buf.WriteString("\n\t\t")
				buf.WriteString("s" + strconv.Itoa(idx))
				buf.WriteString(",")
				idx++
			}
		}

		params := d.Type.Params
		body := d.Body.List[0].(*ast.ReturnStmt).Results[0].(*ast.CompositeLit)
		var l1, l2 token.Pos
		if idx == 0 {
			l1 = params.Opening
			l2 = body.Lbrace
		} else if idx == len(params.List)-1 {
			buf.Write(src[:params.List[len(params.List)-2].End()])
			buf.Write(src[params.List[len(params.List)-1].End():body.Elts[len(body.Elts)-2].End()])
			buf.Write(src[body.Elts[len(body.Elts)-1].End():])
			return
		} else {
			l1 = params.List[idx-1].End()
			l2 = body.Elts[idx-1].End()
		}

		buf.Write(src[:l1])
		writeParams(idx, params.List[idx+1:])
		buf.WriteString("\n")
		buf.Write(src[params.Closing-1 : l2])
		writeBody(idx, len(body.Elts[idx+1:]))
		buf.Write(src[body.Elts[len(body.Elts)-1].End():])
	})

	if _, ok := repeats[path]; m.err == nil && !ok {
		m.pg.Output("\x1b[1A\x1b[1;33m[%10s]\x1b[0m remove service from \x1b[1;34m[%s]\x1b[0m", m.dur(time.Since(before)), path)
		repeats[path] = struct{}{}
	}
}

func (m *Maker) MergeService(path, key string, kind ast.ObjKind, insert func(int) ([]string, []string, error)) {
	if m.err != nil {
		return
	}

	before := time.Now()
	if _, ok := repeats[path]; !ok {
		m.pg.Output("\x1b[1;33m[\x1b[1;5;33m%s\x1b[0m\x1b[1;33m]\x1b[0m merge service to \x1b[1;4;34m[%s]\x1b[0m", placeholder, path)
		m.out = true
	}

	m.parse(path, key, kind, func(src []byte, buf *bytes.Buffer, expr *ast.CallExpr) {
		l := len(expr.Args)
		if l == 0 {
			return
		}

		params, _, err := insert(l)
		if err != nil {
			m.err = err
			return
		}

		if len(expr.Args) == 0 {
			buf.Write(src[:expr.Lparen])
			for _, p := range params {
				buf.WriteString("\n\t")
				buf.WriteString(p)
				buf.WriteString(",")
			}
			buf.WriteString("\n")
			buf.Write(src[expr.Rparen-1:])
			return
		}

		idx := expr.Args[len(expr.Args)-1].End()
		buf.Write(src[:idx])
		for _, p := range params {
			buf.WriteString("\n\t")
			buf.WriteString(p)
			buf.WriteString(",")
		}
		buf.Write(src[idx:])
	}, func(src []byte, buf *bytes.Buffer, d *ast.FuncDecl) {
		params := d.Type.Params
		l := len(params.List)
		if l == 0 {
			return
		}

		body := d.Body.List[0].(*ast.ReturnStmt).Results[0].(*ast.CompositeLit)
		ps, rs, err := insert(l)
		if err != nil {
			m.err = err
			return
		}

		if len(params.List) == 0 {
			buf.Write(src[:params.Opening])
			for _, p := range ps {
				buf.WriteString("\n\t")
				buf.WriteString(p)
				buf.WriteString(",")
			}
			buf.WriteString("\n")
			buf.Write(src[params.Closing-1 : body.Lbrace])
			for _, r := range rs {
				buf.WriteString("\n\t\t")
				buf.WriteString(r)
				buf.WriteString(",")
			}
			buf.WriteString("\n\t")
			buf.Write(src[body.Rbrace-1:])
			return
		}

		idx1 := params.List[len(params.List)-1].Type.End()
		idx2 := body.Elts[len(body.Elts)-1].End()
		buf.Write(src[:idx1])
		for _, p := range ps {
			buf.WriteString("\n\t")
			buf.WriteString(p)
			buf.WriteString(",")
		}
		buf.Write(src[idx1:idx2])
		for _, p := range rs {
			buf.WriteString("\n\t\t")
			buf.WriteString(p)
			buf.WriteString(",")
		}
		buf.Write(src[idx2:])
	})

	if _, ok := repeats[path]; m.err == nil && !ok {
		m.pg.Output("\x1b[1A\x1b[1;33m[%10s]\x1b[0m merge service to \x1b[1;34m[%s]\x1b[0m", m.dur(time.Since(before)), path)
		repeats[path] = struct{}{}
	}
}

func (m *Maker) Chdir(dir string) {
	if m.err != nil {
		return
	}

	m.err = os.Chdir(dir)
}

func (m *Maker) Cmd(command string, args ...string) {
	if m.err != nil {
		return
	}

	before := time.Now()
	names := strings.Split(command, " ")
	if len(args) > 0 {
		command += " " + strings.Join(args, " ")
	}
	lipb := m.libp
	if !strings.HasPrefix(command, "go") {
		lipb = ""
	}
	m.pg.Output("\x1b[1;33m[\x1b[1;5;33m%s\x1b[0m\x1b[1;33m]\x1b[0m %s\x1b[1;4;36m%s\x1b[0m ⏎ ", placeholder, lipb, command)
	m.out = true
	cmd := exec.Command(names[0], append(names[1:], args...)...)
	var buf []byte
	buf, m.err = cmd.CombinedOutput()
	if m.err != nil {
		rtn := string(buf)
		idx := strings.Index(rtn, "exit status")
		if idx > 0 {
			rtn = rtn[:idx]
		}
		rtn = strings.TrimSuffix(rtn, "\n")
		m.err = errors.New(rtn)
		return
	}

	m.pg.Output("\x1b[1A\x1b[1;33m[%10s]\x1b[0m %s\x1b[1;4;36m%s\x1b[0m ⏎ ", m.dur(time.Since(before)), lipb, command)
	m.pg.Output(string(buf))
}

func (m *Maker) Template(path, name string, v any) {
	if m.err != nil {
		return
	}

	m.pg.Output("\x1b[1;33m[\x1b[1;5;33m%s\x1b[0m\x1b[1;33m]\x1b[0m generate go file \x1b[1;4;34m[%s]\x1b[0m", placeholder, path)
	m.out = true
	before := time.Now()
	var f *os.File
	f, m.err = os.Create(path)
	if m.err != nil {
		return
	}
	m.err = m.tp.ExecuteTemplate(f, name, v)
	if m.err != nil {
		return
	}
	m.pg.Output("\x1b[1A\x1b[1;33m[%10s]\x1b[0m generate go file \x1b[1;34m[%s]\x1b[0m", m.dur(time.Since(before)), path)
}

func (m *Maker) MKDir(path string) {
	if m.err != nil {
		return
	}

	m.pg.Output("\x1b[1;33m[\x1b[1;5;33m%s\x1b[0m\x1b[1;33m]\x1b[0m create directory \x1b[1;4;32m[%s]\x1b[0m", placeholder, path)
	m.out = true
	before := time.Now()
	m.err = os.MkdirAll(path, 0744)
	if m.err == nil {
		m.pg.Output("\x1b[1A\x1b[1;33m[%10s]\x1b[0m create directory \x1b[1;32m[%s]\x1b[0m", m.dur(time.Since(before)), path)
	}
}

func (m *Maker) dur(d time.Duration) string {
	ds := d.String()
	if len(ds) > timeLength {
		if strings.HasSuffix(ds, "µs") || strings.HasSuffix(ds, "ms") {
			return strings.TrimSuffix(ds[:timeLength-2], ".") + ds[len(ds)-2:]
		} else if strings.HasSuffix(ds, "s") || strings.HasSuffix(ds, "m") {
			return strings.TrimSuffix(ds[:timeLength-1], ".") + ds[len(ds)-1:]
		}
	}
	return ds
}
