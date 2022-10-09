package configs

import "github.com/gookit/goutil/strutil"

type Environment string

const (
	DEV  Environment = "dev"
	PROD Environment = "prod"
)

type IConfig interface {
	Addr() string
	Env() Environment
	DB() DBConfig
	Redis() RedisConfig
}

type Bootstrap struct {
	Environment Environment `json:"environment,omitempty" yaml:"environment,omitempty"`
	Address     string      `json:"address,omitempty" yaml:"address,omitempty"`
	Database    DBConfig    `json:"database,omitempty" yaml:"database,omitempty"`
	RedisConfig RedisConfig `json:"redis,omitempty" yaml:"redis,omitempty"`
}

func (b Bootstrap) Addr() string {
	return b.Address
}

func (b Bootstrap) Env() Environment {
	if strutil.IsBlank(string(b.Environment)) {
		b.Environment = PROD
	}
	return b.Environment
}

func (b Bootstrap) DB() DBConfig {
	return b.Database
}

func (b Bootstrap) Redis() RedisConfig {
	return b.RedisConfig
}
