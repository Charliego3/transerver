package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"path/filepath"
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
	g.maker.MKDir(g.entPath)
	g.maker.MKDir(g.servicePath)
}

func (g *Generator) genFile() {
	g.maker.Template(filepath.Join(g.cmdPath, "main.go"), "main.go", nil)
	g.maker.Template(filepath.Join(g.cmdPath, "wire.go"), "wire.go", nil)
	g.maker.Template(filepath.Join(g.bizPath, "biz.go"), "biz.go", nil)
	g.maker.Template(filepath.Join(g.confPath, "conf.go"), "conf.go", nil)
	g.maker.Template(filepath.Join(g.dataPath, "data.go"), "data.go", nil)
	g.maker.Template(filepath.Join(g.servicePath, "service.go"), "service.go", nil)

	g.maker.Template(filepath.Join(g.bizPath, "greeter.go"), "bgreeter.go", nil)
	g.maker.Template(filepath.Join(g.dataPath, "greeter.go"), "dgreeter.go", nil)
	g.maker.Template(filepath.Join(g.servicePath, "greeter.go"), "sgreeter.go", nil)
}

func (g *Generator) runCmd() {
	g.maker.Chdir(g.cfg.modPath)
	g.maker.Cmd("go", "mod", "init", g.cfg.modURL)
	g.maker.Cmd("go", "get", "-u", "entgo.io/ent/cmd/ent")
	g.maker.Cmd("go", "get", "-u", "github.com/google/wire")
	g.maker.Chdir(g.internalPath)
	g.maker.Cmd("go", append([]string{"run", "entgo.io/ent/cmd/ent", "init"}, g.cfg.services...)...)
	if g.usingWork() {
		g.maker.Chdir(filepath.Dir(g.cfg.modPath))
		g.maker.Cmd("go", "work", "use", g.cfg.modName)
		g.maker.Cmd("go", "work", "sync")
	}
	g.maker.Chdir(g.entPath)
	g.maker.Cmd("go", "run", "entgo.io/ent/cmd/ent", "generate", "./schema")
	g.maker.Chdir(g.cmdPath)
	g.maker.Cmd("go", "run", "github.com/google/wire/cmd/wire", "./...")
}

func (g *Generator) usingWork() bool {
	_, err := os.Stat(filepath.Join(filepath.Dir(g.cfg.modPath), "go.work"))
	return err == nil
}

func (g *Generator) gen() {
	g.pg.NewLine()
	g.solvePath()
	switch g.cfg.typ {
	case 0: // Create Module
		g.genFile()
		g.runCmd()
	case 1: // Add Service
	case 2: // Remove Module
	case 3: // Remove Service
	}

	if g.err == nil {
		g.err = g.maker.err
	}

	if g.err != nil {
		g.pg.Output(ExitErrStyle.Render(g.err.Error() + "\nexit status 1"))
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
