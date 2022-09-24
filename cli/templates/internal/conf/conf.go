package conf

import (
	"github.com/google/wire"
	"github.com/transerver/commons/configs"
)

var ProviderSet = wire.NewSet(
	NewBootstrap,
)

type Config struct {
	Addr string
}

func NewBootstrap() any {
	return &Config{}
}

func (c Config) Address() string {
	return c.Addr
}

func (c Config) Env() configs.Environment {
	return configs.DEV
}
