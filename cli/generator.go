package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gookit/goutil/strutil"
	"go/ast"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Generator struct {
	spinner spinner.Model
	cfg     *Config
	pg      *Program
	maker   *Maker
	done    bool

	cmdPath      string
	internalPath string
	bizPath      string
	confPath     string
	dataPath     string
	entPath      string
	servicePath  string

	err error
}

func (g *Generator) solvePath() {
	g.cmdPath = filepath.Join(g.cfg.modPath, "cmd")
	g.internalPath = filepath.Join(g.cfg.modPath, "internal")
	g.bizPath = filepath.Join(g.internalPath, "biz")
	g.confPath = filepath.Join(g.internalPath, "conf")
	g.dataPath = filepath.Join(g.internalPath, "data")
	g.entPath = filepath.Join(g.internalPath, "ent")
	g.servicePath = filepath.Join(g.internalPath, "service")

	if g.cfg.typ > 0 {
		return
	}

	g.maker.MKDir(g.cfg.modPath)
	g.maker.MKDir(g.cmdPath)
	g.maker.MKDir(g.internalPath)
	g.maker.MKDir(g.bizPath)
	g.maker.MKDir(g.confPath)
	g.maker.MKDir(g.dataPath)
	g.maker.MKDir(g.servicePath)
}

func (g *Generator) genBizs() {
	if g.err != nil {
		return
	}

	g.maker.Template(filepath.Join(g.cmdPath, "main.go"), "main.gohtml", g.cfg)
	g.maker.Template(filepath.Join(g.cmdPath, "wire.go"), "wire.gohtml", g.cfg)
	g.maker.Template(filepath.Join(g.bizPath, "biz.go"), "biz.gohtml", g.cfg)
	g.maker.Template(filepath.Join(g.confPath, "conf.go"), "conf.gohtml", nil)
	g.maker.Template(filepath.Join(g.confPath, "config.yaml"), "config.yaml", nil)
	g.maker.Template(filepath.Join(g.dataPath, "data.go"), "data.gohtml", g.cfg)
	g.maker.Template(filepath.Join(g.servicePath, "service.go"), "service.gohtml", g.cfg)
}

func (g *Generator) genServices() {
	if g.err != nil {
		return
	}

	for _, name := range g.cfg.Services {
		lowerName := strings.ToLower(name)
		fileName := lowerName + ".go"
		g.cfg.CurrService = name
		g.cfg.CurrServiceLower = lowerName
		g.maker.Template(filepath.Join(g.bizPath, fileName), "bgreeter.gohtml", g.cfg)
		g.maker.Template(filepath.Join(g.dataPath, fileName), "dgreeter.gohtml", g.cfg)
		g.maker.Template(filepath.Join(g.servicePath, fileName), "sgreeter.gohtml", g.cfg)
	}
}

func (g *Generator) createModule() {
	if g.err != nil {
		return
	}

	g.maker.Chdir(g.cfg.modPath)
	g.maker.Cmd("go mod init", g.cfg.ModURL)
	if g.usingWork() {
		g.maker.Chdir(filepath.Dir(g.cfg.modPath))
		g.maker.Cmd("go work use", g.cfg.ModName)
		g.maker.Cmd("go work sync")
	}
	g.genEnt()
	g.maker.Chdir(g.cmdPath)
	g.maker.Cmd("go get -u github.com/google/wire")
	g.maker.Cmd("go run github.com/google/wire/cmd/wire ./...")
}

func (g *Generator) genEnt() {
	if g.err != nil || len(g.cfg.Services) == 0 {
		return
	}

	g.maker.Chdir(g.internalPath)
	g.maker.Cmd("go get -u entgo.io/ent/cmd/ent")
	g.maker.Cmd("go run entgo.io/ent/cmd/ent init", g.cfg.Services...)
	g.maker.Chdir(g.entPath)
	g.maker.Cmd("go run entgo.io/ent/cmd/ent generate ./schema")
}

func (g *Generator) usingWork() bool {
	_, err := os.Stat(filepath.Join(filepath.Dir(g.cfg.modPath), "go.work"))
	return err == nil
}

func (g *Generator) addService(suffix string) func(int) ([]byte, []byte, error) {
	return func(int) ([]byte, []byte, error) {
		var buf bytes.Buffer
		for _, s := range g.cfg.Services {
			buf.WriteString(",\n\tNew" + s + suffix)
		}
		return buf.Bytes(), nil, nil
	}
}

func (g *Generator) margeService() {
	spath := filepath.Join(g.servicePath, "service.go")
	g.maker.MergeService(filepath.Join(g.bizPath, "biz.go"), "ProviderSet", ast.Var, g.addService("Usecase"))
	g.maker.MergeService(filepath.Join(g.dataPath, "data.go"), "ProviderSet", ast.Var, g.addService("Repo"))
	g.maker.MergeService(spath, "ProviderSet", ast.Var, g.addService("Service"))
	g.maker.MergeService(spath, "MakeServices", ast.Fun, func(count int) ([]byte, []byte, error) {
		var p, r bytes.Buffer
		for _, s := range g.cfg.Services {
			p.WriteString(",\n\ts" + strconv.Itoa(count) + " *" + s + "Service")
			r.WriteString(",\n\t\ts" + strconv.Itoa(count))
			count++
		}
		return p.Bytes(), r.Bytes(), nil
	})
}

func (g *Generator) revertService() {

}

func (g *Generator) revertWork() {

}

func (g *Generator) restoreModURL() {
	modPath := filepath.Join(g.cfg.modPath, "go.mod")
	var f *os.File
	f, g.err = os.Open(modPath)
	if g.err != nil {
		return
	}

	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		g.err = fmt.Errorf("go.mod mayby empty")
		return
	}

	txt := scanner.Text()
	var ts []string
	if strutil.IsNotBlank(txt) {
		ts = strings.Split(txt, " ")
	}
	if len(ts) < 2 {
		g.err = fmt.Errorf("go.mod is invalid")
		return
	}
	g.cfg.ModURL = ts[1]
}

func (g *Generator) gen() {
	g.pg.NewLine()
	g.solvePath()
	switch g.cfg.typ {
	case 0: // Create Module
		g.genBizs()
		g.genServices()
		g.createModule()
	case 1: // Add Service
		g.restoreModURL()
		g.genEnt()
		g.genServices()
		g.margeService()
	case 2: // Remove Module
		// remove dir // TODO
		g.revertWork() // TODO
	case 3: // Remove Service
		// remove service // TODO
		g.revertService() // TODO
	}

	if g.err == nil {
		g.err = g.maker.err
	}

	if g.err != nil {
		msg := ExitErrStyle.Render(g.err.Error() + "\nexit status 1")
		if !g.maker.out {
			msg = "\x1b[1A" + msg
		}
		g.pg.Output(msg)
	}

	g.done = true
}

func NewGenerator(cfg *Config, pg *Program) *Generator {
	s := spinner.New()
	s.Spinner = spinner.Points
	s.Style = SpinnerStyle
	g := &Generator{spinner: s, cfg: cfg, pg: pg, maker: NewMaker(pg)}
	go g.gen()
	return g
}

func (g *Generator) Update(msg tea.Msg) tea.Cmd {
	if g.done {
		return tea.Quit
	}

	var cmd tea.Cmd
	g.spinner, cmd = g.spinner.Update(msg)
	return cmd
}

func (g *Generator) View() string {
	if g.done {
		return "\n"
	}
	return TLBmarginStyle.Render(g.spinner.View()) + BlurredStyle.Render(" waiting for generate...\n")
}

func (g *Generator) Callback(*Program) (string, bool) {
	return "", false
}
