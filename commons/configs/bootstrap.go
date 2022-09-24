package configs

type Environment uint

const (
	DEV Environment = iota + 1
	PROD
)

type Bootstrap interface {
	Address() string
	Env() Environment
}
