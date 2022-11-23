package configs

import (
	"github.com/charliego93/flagx"
)

func init() {
	flagx.SetErrorHandling(flagx.SkipNoDeclared | flagx.ClearAfterParse)
	cfgPath := flagx.String("config,c", "config.yaml", flagx.WithDescription("config file path"))
	err := flagx.Parse()
	if err != nil {
		panic(err)
	}

	root = &Base{}
	Bootstrap = root
	Parse(NewFileLoader(*cfgPath))
}
