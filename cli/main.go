package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/gookit/goutil/strutil"
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
	typ            int // actions index
	modName        string
	modPath        string
	modURL         string
	usingCfgOpt    bool
	usingETCDOpt   bool
	usingGRPCOpt   bool
	hsOpt          int
	usingLogOpt    bool
	usingRedisOpts bool
	services       []string
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
		pg.AddModel(askModURL)
		pg.AddModel(selectHsOpts)
		pg.AddModel(askCfgOpts)
		pg.AddModel(askETCDOpts)
		pg.AddModel(askGRPCOpts)
		pg.AddModel(askLoggerOpts)
		pg.AddModel(askRedisOpts)
		pg.AddModel(askAddService)
	case 1: // Add Service
		pg.AddModel(askModPath)
		pg.AddModel(askAddService)
		pg.AddModel(generate)
	case 2: // Remove Module
		pg.AddModel(askModPath)
		pg.AddModel(generate)
	case 3: // remove Service
		pg.AddModel(askModPath)
		pg.AddModel(inputServiceNames)
		pg.AddModel(generate)
	}
	return output("Generate Type", m.SelectedItem().FilterValue()), false
}

func generate() IModel {
	return NewGenerator(cfg, pg)
}

func askModURL() IModel {
	return NewInput("What is mod init URL?",
		WithPlaceholder(cfg.modName),
		WithInputCallback(func(url string) (string, bool) {
			if strutil.IsBlank(url) {
				url = cfg.modName
			}
			cfg.modURL = url
			return output("Mod init URL", cfg.modURL), false
		}))
}

func askAddService() IModel {
	return NewConfirm("Want to add service?", func(b bool) (string, bool) {
		if b {
			pg.AddModel(inputServiceNames)
		}
		pg.AddModel(generate)
		return "", false
	})
}

func inputServiceNames() IModel {
	return NewInput("Please enter service name", WithMulti(),
		WithInputCallbacks(func(s []string) (string, bool) {
			if len(s) == 0 {
				cfg.services = []string{"Greeter"}
			} else {
				cfg.services = s
			}
			return output("Service names", "[%s]", strings.Join(s, ", ")), false
		}))
}

func askRedisOpts() IModel {
	return NewConfirm("Want using redis options?", func(b bool) (string, bool) {
		cfg.usingRedisOpts = b
		return boutput("Using redis options", b), false
	})
}

func askLoggerOpts() IModel {
	return NewConfirm("Want using logger writer?", func(b bool) (string, bool) {
		cfg.usingLogOpt = b
		return boutput("Using logger writer", b), false
	})
}

func selectHsOpts() IModel {
	items := []list.Item{
		SelectItem("Using full options"),
		SelectItem("Without ServeMuxOption but using handlers"),
		SelectItem("Without any options"),
	}
	return NewSelect(items,
		WithTitle("What type using HTTP server options?"),
		WithShowHelp(),
		WithCallback(func(m list.Model) (string, bool) {
			cfg.hsOpt = m.Index()
			return output("Using HTTP options", m.SelectedItem().FilterValue()), false
		}))
}

func askETCDOpts() IModel {
	return NewConfirm("Want using ETCD optoins?", func(b bool) (string, bool) {
		cfg.usingETCDOpt = b
		return boutput("Using ETCD options", b), false
	})
}

func askCfgOpts() IModel {
	return NewConfirm("Want using Config parser optoins?", func(b bool) (string, bool) {
		cfg.usingCfgOpt = b
		return boutput("Using Config parser options", b), false
	})
}

func askModName() IModel {
	return NewInput(
		"What is the module name?",
		WithPlaceholder("please enter a module name"),
		WithInputCallback(func(s string) (string, bool) {
			if strutil.IsBlank(s) {
				return Exit("Not input anything for mod name..."), true
			}
			cfg.modName = s

			if cfg.modName == "g" {
				_ = os.RemoveAll("g")
			}
			return output("Module name", cfg.modName), false
		}))
}

func askModPath() IModel {
	return NewInput(
		"What is the mod directory?",
		WithPlaceholder(filepath.Join(wd, cfg.modName)),
		WithInputCallback(func(path string) (string, bool) {
			if strutil.IsBlank(path) {
				if cfg.typ != 0 {
					return Exit("Please enter absolute path or module name"), true
				}
				cfg.modPath = filepath.Join(wd, cfg.modName)
			} else {
				if filepath.IsAbs(path) {
					cfg.modPath = path
				} else {
					cfg.modPath = filepath.Join(wd, path)
				}

				if cfg.typ > 0 {
					cfg.modName = filepath.Base(cfg.modPath)
				}
			}

			_, err := os.Stat(cfg.modPath)
			if cfg.typ == 0 && err == nil {
				return Exit("Directory is already exits: [%s]", cfg.modPath), true
			} else if cfg.typ > 0 && err != nil {
				return Exit("Directory is not exits: [%s]", cfg.modPath), true
			}
			return output("Directory path", cfg.modPath), false
		}))
}

func askGRPCOpts() IModel {
	return NewConfirm("Want using GRPC optoins?", func(b bool) (string, bool) {
		cfg.usingGRPCOpt = b
		return boutput("Using GRPC options", b), false
	})
}

func Exit(format string, v ...any) string {
	return ExitErrStyle.Render(fmt.Sprintf(format+"\nexit status 1", v...))
}

func boutput(key string, b bool) string {
	return output(key, strconv.FormatBool(b))
}

func output(key, format string, v ...any) string {
	var result string
	if !marginTop {
		result += "\n"
		marginTop = true
	}
	if format == "true" || format == "false" {
		format = strings.ToUpper(format)
	}
	return result + ResultKeyStyle.Render(key+":") + " " + ResultStyle.Render(fmt.Sprintf(format, v...))
}
