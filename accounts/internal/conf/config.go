package conf

import (
	"github.com/transerver/commons/configs"
)

var (
	// check type
	_ configs.IBootstrap = (*bootstrap)(nil)

	// Bootstrap instance
	Bootstrap = &bootstrap{}
)

type bootstrap struct {
	configs.Base `json:",inline" yaml:",inline"`
}

func init() {
	configs.RegisterBootstrap(Bootstrap)
	configs.Parse(configs.NewFileLoader(*path))
}
