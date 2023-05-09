package grpcx

import (
	"net"

	"github.com/transerver/app/logger"
	"github.com/transerver/app/opts"
	"google.golang.org/grpc"
)

// WithAddr create a listener with network and address
// WithAddr and WithListener just choose one of them
func WithAddr(network, addr string) opts.Option[Server] {
	return opts.OptionFunc[Server](func(cfg *Server) {
		listenr, err := net.Listen(network, addr)
		if err != nil {
			logger.Fatal("failed to listen grpc server", "err", err)
		}
		cfg.listener = listenr
	})
}

// WithListener uses the given listener
// WithListener and WithAddr just choose one of them
func WithListener(lis net.Listener) opts.Option[Server] {
	return opts.OptionFunc[Server](func(cfg *Server) {
		cfg.listener = lis
	})
}

// WithServerOption inject grpc.ServerOption to server
func WithServerOption(gsos ...grpc.ServerOption) opts.Option[Server] {
	return opts.OptionFunc[Server](func(cfg *Server) {
		cfg.srvOpts = gsos
	})
}
