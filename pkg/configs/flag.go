package configs

import (
	"github.com/charliego93/flagx"
)

var cfgPath = flagx.String("config,c", "config.yaml", flagx.WithDescription("config file path"))

func init() {
	flagx.SetErrorHandling(flagx.SkipNoDeclared | flagx.OverrideRedefined)
	err := flagx.Parse()
	if err != nil {
		panic(err)
	}

	root = &Base{}
	Bootstrap = root
	Parse(NewFileLoader(*cfgPath))
}
