package redis

import (
	"time"

	rv9 "github.com/go-redis/redis/v9"
)

type Client struct {
	rv9.UniversalClient
}

// NewClientWithoutOpts returns redis UniversalClient wrapper
// using all of default prop
func NewClientWithoutOpts() *Client {
	return NewClient()
}

// NewClient returns redis UniversalClient wrapper
// if the options is empty using default address with 127.0.0.1:6379
// if not specified DialTimeout default is a minute
func NewClient(opts ...Option) *Client {
	cfg := &rv9.UniversalOptions{}
	for _, opt := range opts {
		opt(cfg)
	}

	if len(cfg.Addrs) == 0 {
		cfg.Addrs = []string{"127.0.0.1:6379"}
	}

	if cfg.DialTimeout == 0 {
		cfg.DialTimeout = time.Minute
	}

	cli := rv9.NewUniversalClient(cfg)
	return &Client{UniversalClient: cli}
}
