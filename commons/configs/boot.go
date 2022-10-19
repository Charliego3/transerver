package configs

import "github.com/gookit/goutil/strutil"

type Environment string

var Bootstrap IBootstrap

func RegisterBootstrap(boot IBootstrap) {
	Bootstrap = boot
}

const (
	DEV  Environment = "dev"
	PROD Environment = "prod"
)

type IBootstrap interface {
	Root() *Base
}

type Base struct {
	Environment Environment `json:"environment,omitempty" yaml:"environment,omitempty"`
	Name        string      `json:"name,omitempty" yaml:"name,omitempty"`
	Network     string      `json:"network,omitempty" yaml:"network,omitempty"`
	Address     string      `json:"address,omitempty" yaml:"address,omitempty"`
	Database    Database    `json:"database,omitempty" yaml:"database,omitempty"`
	Redis       Redis       `json:"redis,omitempty" yaml:"redis,omitempty"`
	Etcd        ETCD        `json:"etcd,omitempty" yaml:"etcd,omitempty"`
}

func (b *Base) Root() *Base {
	return b
}

func (b *Base) Env() Environment {
	if strutil.IsBlank(string(b.Environment)) {
		return PROD
	}
	return b.Environment
}
