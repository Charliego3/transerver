package mapp

import (
	"net"

	"github.com/transerver/mapp/grpcx"
	"github.com/transerver/mapp/httpx"
	"google.golang.org/grpc"

	"github.com/charliego93/logger"
	"github.com/transerver/mapp/opts"
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

	// disableHTTP only serve grpc server
	disableHTTP bool

	// disableGRPC only serve http server
	disableGRPC bool

	// onStartup run on Applition after init
	onStartup func(*Application) error

	// gopts is grpcx.Server options
	gopts []opts.Option[grpcx.Server]

	// middles accept http server Middleware
	hopts []opts.Option[httpx.Server]

	logger logger.Logger
}

func DisableHTTP() opts.Option[Config] {
	return opts.OptionFunc[Config](func(cfg *Config) {
		cfg.disableHTTP = true
	})
}

func DisableGRPC() opts.Option[Config] {
	return opts.OptionFunc[Config](func(cfg *Config) {
		cfg.disableGRPC = true
	})
}

func WithHttpOpts(hopts ...opts.Option[httpx.Server]) opts.Option[Config] {
	return opts.OptionFunc[Config](func(cfg *Config) {
		cfg.hopts = hopts
	})
}

func OnStartup(fn func(*Application) error) opts.Option[Config] {
	return opts.OptionFunc[Config](func(cfg *Config) {
		cfg.onStartup = fn
	})
}

// WithGrpcServerOpts accept grpc server options
func WithGrpcServerOpts(gopts ...grpc.ServerOption) opts.Option[Config] {
	return opts.OptionFunc[Config](func(cfg *Config) {
		cfg.gopts = append(cfg.gopts, grpcx.WithServerOption(gopts...))
	})
}

// WithAddr served http and grpc on same address
func WithAddr(network, addr string) opts.Option[Config] {
	return opts.OptionFunc[Config](func(cfg *Config) {
		listener, err := net.Listen(network, addr)
		if err != nil {
			cfg.logger.Fatal("failed to listen app", "err", err)
		}
		cfg.lis = listener
	})
}

// WithTCPAddr is WithAddr alias but network using tcp
func WithTCPAddr(addr string) opts.Option[Config] {
	return WithAddr("tcp", addr)
}

// WithListener served http and grpc on same address
func WithListener(lis net.Listener) opts.Option[Config] {
	return opts.OptionFunc[Config](func(cfg *Config) {
		cfg.lis = lis
	})
}

// WithHttpAddr expected http server listen address using tcp network
func WithHttpAddr(network, addr string) opts.Option[Config] {
	return opts.OptionFunc[Config](func(cfg *Config) {
		listener, err := net.Listen(network, addr)
		if err != nil {
			cfg.logger.Fatal("failed to listen http server with app", "err", err)
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
			cfg.logger.Fatal("failed to listen grpc server with app", "err", err)
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
