package redis

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	rv9 "github.com/go-redis/redis/v9"
)

type Option func(opt *rv9.UniversalOptions)

func WithAddrs(addrs ...string) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.Addrs = addrs
	}
}

func WithDialer(dialer func(context.Context, string, string) (net.Conn, error)) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.Dialer = dialer
	}
}

func WithOnConnect(fn func(context.Context, *rv9.Conn) error) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.OnConnect = fn
	}
}

func WithUsername(username string) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.Username = username
	}
}

func WithPassword(password string) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.Password = password
	}
}

func WithSentinelUsername(username string) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.SentinelUsername = username
	}
}

func WithSentinelPassword(password string) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.SentinelPassword = password
	}
}

func WithMaxRetries(max int) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.MaxRetries = max
	}
}

func WithMinRetryBackoff(min time.Duration) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.MinRetryBackoff = min
	}
}

func WithMaxRetryBackoff(max time.Duration) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.MaxRetryBackoff = max
	}
}

func WithDialTimeout(timeout time.Duration) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.DialTimeout = timeout
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.ReadTimeout = timeout
	}
}

func WithWriteTimout(timeout time.Duration) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.WriteTimeout = timeout
	}
}

func WithPoolFIFO(fifo bool) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.PoolFIFO = fifo
	}
}

func WithPoolSize(size int) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.PoolSize = size
	}
}

func WithPoolTimeout(timeout time.Duration) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.PoolTimeout = timeout
	}
}

func WithMinIdleConns(conns int) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.MinIdleConns = conns
	}
}

func WithMaxIdleConns(conns int) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.MaxIdleConns = conns
	}
}

func WithConnMaxIdleTime(max time.Duration) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.ConnMaxIdleTime = max
	}
}

func WithConnMaxLifetime(lifetime time.Duration) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.ConnMaxLifetime = lifetime
	}
}

func WithTLSConfig(tls *tls.Config) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.TLSConfig = tls
	}
}

func WithMaxRedirects(max int) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.MaxRedirects = max
	}
}

func WithReadOnly(readOnly bool) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.ReadOnly = readOnly
	}
}

func WithRouteByLatency(latency bool) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.RouteByLatency = latency
	}
}

func WithRouteRandomly(randomly bool) Option {
	return func(opt *rv9.UniversalOptions) {
		opt.RouteRandomly = randomly
	}
}
