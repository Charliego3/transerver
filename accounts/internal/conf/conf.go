package conf

import (
	"github.com/google/wire"
	"github.com/transerver/commons/configs"
)

var _ configs.IConfig = (*Bootstrap)(nil)
var ProviderSet = wire.NewSet(
	NewBootstrap,
)

type Bootstrap struct {
	configs.Bootstrap `json:",inline" yaml:",inline"`
}

func NewBootstrap() any {
	return &Bootstrap{}
}
