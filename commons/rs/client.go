package rs

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	rv9 "github.com/go-redis/redis/v9"
)

type Config struct {
	Dialer    func(ctx context.Context, network, addr string) (net.Conn, error)
	OnConnect func(ctx context.Context, cn *rv9.Conn) error
	TLSConfig *tls.Config
}

type Client struct {
	rv9.UniversalClient
}

// New returns redis UniversalClient wrapper
// if the options is empty using default address with 127.0.0.1:6379
// if not specified DialTimeout default is a minute
func New(opts ...Option) (*Client, error) {
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

	timeout := cfg.ReadTimeout
	if timeout == 0 {
		timeout = cfg.DialTimeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &Client{UniversalClient: cli}, nil
}
