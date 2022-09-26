package configs

type Environment uint

const (
	DEV Environment = iota + 1
	PROD
)

type IConfig interface {
	Addr() string
	Env() Environment
	DB() Database
}

type Bootstrap struct {
	Environment Environment `json:"environment,omitempty" yaml:"environment,omitempty"`
	Address     string      `json:"address,omitempty" yaml:"address,omitempty"`
	Database    Database    `json:"database,omitempty" yaml:"database,omitempty"`
}

func (b Bootstrap) Addr() string {
	return b.Address
}

func (b Bootstrap) Env() Environment {
	return b.Environment
}

func (b Bootstrap) DB() Database {
	return b.Database
}
