package app

import (
	"net"

	"github.com/transerver/app/logger"
	"github.com/transerver/app/opts"
)

type Config struct {
	// http and grpc server both listen this address
	lis net.Listener

	// http server to serve this address
	// if lis and hlis both nil, then hlis using dynamic address
	hlis net.Listener

	// grpc server to serve this address
	// if lis and glis both nil, then glis using dynamic address
	glis net.Listener
}

// WithAddr served http and grpc on same address
func WithAddr(network, addr string) opts.Option[Config] {
	return opts.OptionFunc[Config](func(cfg *Config) {
		listener, err := net.Listen(network, addr)
		if err != nil {
			logger.Fatal("app listen fail", "err", err)
		}
		cfg.lis = listener
	})
}

// WithListener served http and grpc on same address
func WithListener(lis net.Listener) opts.Option[Config] {
	return opts.OptionFunc[Config](func(cfg *Config) {
		cfg.lis = lis
	})
}

// WithHttpAddr execpted http server listen address using tcp network
func WithHttpAddr(network, addr string) opts.Option[Config] {
	return opts.OptionFunc[Config](func(cfg *Config) {
		listener, err := net.Listen(network, addr)
		if err != nil {
			logger.Fatal("http server listen fail", "err", err)
		}
		cfg.hlis = listener
	})
}

// WithHttpListener served http server listener
func WithHttpListener(lis net.Listener) opts.Option[Config] {
	return opts.OptionFunc[Config](func(cfg *Config) {
		cfg.hlis = lis
	})
}

// WithGrpcAddr served grpc server on address
func WithGrpcAddr(network, addr string) opts.Option[Config] {
	return opts.OptionFunc[Config](func(cfg *Config) {
		listener, err := net.Listen(network, addr)
		if err != nil {
			logger.Fatal("grpc server listen fail", "err", err)
		}
		cfg.glis = listener
	})
}

// WithGrpcListener served grpc server listener
func WithGrpcListener(lis net.Listener) opts.Option[Config] {
	return opts.OptionFunc[Config](func(cfg *Config) {
		cfg.glis = lis
	})
}
