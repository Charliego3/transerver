package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"github.com/gookit/goutil/strutil"
	"os"
	"path/filepath"
	"strconv"
)

var (
	marginTop bool
	pg        = NewProgram()
	cfg       = &Config{}
	wd, _     = os.Getwd()
	actions   = []list.Item{
		SelectItem("Create Module"),
		SelectItem("Add Service"),
		SelectItem("Remove Module"),
		SelectItem("Remove Service"),
	}
)

type Config struct {
	typ     int // actions index
	modName string
	modPath string
}

func main() {
	pg.AddModel(func() IModel {
		return NewSelect(
			actions, WithTitle("What is want to do generate?"),
			WithShowHelp(), WithCallback(ActionsCallback))
	})

	pg.Start()
}

func ActionsCallback(m list.Model) (string, bool) {
	cfg.typ = m.Index()
	switch cfg.typ {
	case 0: // Create Module
		pg.AddModel(askModName)
		pg.AddModel(askModPath)
	case 1: // Add Service
		pg.AddModel(askGRPCOpts)
	case 2: // Remove Module
	case 3: // remove Service

	}
	return output("Select mode", string(m.SelectedItem().(SelectItem))), false
}

func askModName() IModel {
	return NewInput(
		"What is the module name?",
		WithPlaceholder("please enter a module name"),
		WithInputCallback(func(s string) (string, bool) {
			if strutil.IsBlank(s) {
				return exit("Not input anything for mod name..."), true
			}
			cfg.modName = s
			return output("Module name", cfg.modName), false
		}))
}

func askModPath() IModel {
	return NewInput(
		"What is the directory?",
		WithPlaceholder(filepath.Join(wd, cfg.modName)),
		WithInputCallback(func(s string) (string, bool) {
			if strutil.IsBlank(s) {
				cfg.modPath = filepath.Join(wd, cfg.modName)
			} else {
				cfg.modPath = s
			}

			_, err := os.Stat(cfg.modPath)
			if err == nil {
				return exit("directory is already exits: [%s]", cfg.modPath), true
			}
			return output("Directory path", cfg.modPath), false
		}))
}

func askGRPCOpts() IModel {
	return NewConfirm("Want using GRPC optoins?", func(b bool) (string, bool) {
		return output("Using GRPC options", strconv.FormatBool(b)), false
	})
}

func exit(format string, v ...any) string {
	return ExitErrStyle.Render(fmt.Sprintf(format+"\nexit status 1", v...))
}

func output(key, format string, v ...any) string {
	var result string
	if !marginTop {
		result += "\n"
		marginTop = true
	}
	return result + ResultKeyStyle.Render(key+":") + " " + ResultStyle.Render(fmt.Sprintf(format, v...))
}
