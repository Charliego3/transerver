package main

import (
	"embed"
	"errors"
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
)

type Maker struct {
	tp   *template.Template
	pg   *Program
	libp string
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

func (m *Maker) Chdir(dir string) {
	if m.err != nil {
		return
	}

	m.err = os.Chdir(dir)
}

func (m *Maker) Cmd(name string, args ...string) {
	if m.err != nil {
		return
	}

	before := time.Now()
	cmd := exec.Command(name, args...)
	command := strings.TrimPrefix(cmd.String(), m.libp)
	m.pg.Output("\x1b[1;33m[\x1b[1;5;33m%s\x1b[0m\x1b[1;33m]\x1b[0m %s\x1b[1;4;36m%s\x1b[0m ⏎ ", placeholder, m.libp, command)

	var buf []byte
	buf, m.err = cmd.CombinedOutput()
	if m.err != nil {
		rtn := string(buf)
		rtn = rtn[:strings.Index(rtn, "exit status")]
		rtn = strings.TrimSuffix(rtn, "\n")
		m.err = errors.New(rtn)
		return
	}

	m.pg.Output("\x1b[1A\x1b[1;33m[%10s]\x1b[0m %s\x1b[1;4;36m%s\x1b[0m ⏎ ", m.dur(time.Since(before)), m.libp, command)
	m.pg.Output(string(buf))
}

func (m *Maker) Template(path, name string, v any) {
	if m.err != nil {
		return
	}

	m.pg.Output("\x1b[1;33m[\x1b[1;5;33m%s\x1b[0m\x1b[1;33m]\x1b[0m generate go file \x1b[1;4;34m[%s]\x1b[0m", placeholder, path)
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

	m.pg.Output("\x1b[1;5;33m[\x1b[1;33m%s\x1b[0m\x1b[1;5;33m]\x1b[0m create directory \x1b[1;4;32m[%s]\x1b[0m", placeholder, path)
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
