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

func (m *Maker) MergeService(path, key string, kind ast.ObjKind, insert func(int) ([]byte, []byte, error)) {
	if m.err != nil {
		return
	}

	before := time.Now()
	var src []byte
	src, m.err = os.ReadFile(path)
	if m.err != nil {
		return
	}

	fset := token.NewFileSet()
	if _, ok := repeats[path]; !ok {
		m.pg.Output("\x1b[1;33m[\x1b[1;5;33m%s\x1b[0m\x1b[1;33m]\x1b[0m merge service to file \x1b[1;4;34m[%s]\x1b[0m", placeholder, path)
		m.out = true
	}
	var f *ast.File
	f, m.err = parser.ParseFile(fset, "", src, parser.DeclarationErrors)
	if m.err != nil {
		return
	}

	for k, obj := range f.Scope.Objects {
		if m.err != nil {
			return
		}

		if k != key {
			continue
		}

		if obj.Kind != kind {
			continue
		}

		var buf bytes.Buffer
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

				params, _, err := insert(len(ce.Args))
				if err != nil {
					m.err = err
					return
				}

				idx := ce.Args[len(ce.Args)-1].End() - 1
				buf.Write(src[:idx])
				buf.Write(params)
				buf.Write(src[idx:])
			}
		case *ast.FuncDecl:
			params := d.Type.Params.List
			bodys := d.Body.List
			pidx := params[len(params)-1].Type.(*ast.StarExpr).X.End() - 1
			arr := bodys[len(bodys)-1].(*ast.ReturnStmt).Results
			elts := arr[len(arr)-1].(*ast.CompositeLit).Elts
			ridx := elts[len(elts)-1].(*ast.Ident).End() - 1

			p, r, err := insert(len(params))
			if err != nil {
				m.err = err
				return
			}

			buf.Write(src[:pidx])
			buf.Write(p)
			buf.Write(src[pidx:ridx])
			buf.Write(r)
			buf.Write(src[ridx:])
		}
		m.err = os.WriteFile(path, buf.Bytes(), 0744)
		if m.err == nil {
			if _, ok := repeats[path]; !ok {
				m.pg.Output("\x1b[1A\x1b[1;33m[%10s]\x1b[0m merge service to file \x1b[1;34m[%s]\x1b[0m", m.dur(time.Since(before)), path)
				repeats[path] = struct{}{}
			}
			before = time.Now()
		}
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
