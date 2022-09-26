package conf

import (
	"github.com/google/wire"
	"github.com/transerver/commons/configs"
)

var ProviderSet = wire.NewSet(
	NewBootstrap,
)

type Config struct {
	Environment configs.Environment `json:"environment,omitempty" yaml:"environment"`
	Addr        string              `json:"addr,omitempty" yaml:"addr"`
	Database    Database            `json:"database" yaml:"database"`
}

func NewBootstrap() any {
	return &Config{}
}

func (c Config) Address() string {
	return c.Addr
}

func (c Config) Env() configs.Environment {
	return c.Environment
}
