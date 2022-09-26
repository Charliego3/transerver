package configs

import "io"

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

func WithReadCloser(r io.ReadCloser) Option {
	return func(c *Config) {
		c.reader = r
	}
}
