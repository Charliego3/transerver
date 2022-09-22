package config

type Option func(*Config)

func WithPath(paths ...string) Option {
	return func(c *Config) {
		c.paths = paths
	}
}

func WithSource(source []byte) Option {
	return func(c *Config) {
		c.source = source
	}
}
