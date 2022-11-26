package configs

import (
	"flag"
	"github.com/charliego93/flagx"
)

var cfgPath = flagx.String("config,c", "config.yaml", flagx.WithDescription("config file path"))
var _ = flag.String("config", "config.yaml", "config usage")

func init() {
	flagx.SetErrorHandling(flagx.SkipNoDeclared | flagx.ClearAfterParse)
	err := flagx.Parse()
	if err != nil {
		panic(err)
	}

	root = &Base{}
	Bootstrap = root
	Parse(NewFileLoader(*cfgPath))
}
