package grpcx

import (
	"github.com/transerver/app/opts"
	"google.golang.org/grpc"
)

func WithAddr(addr string) opts.Option[Server] {
	return opts.OptionFunc[Server](func(cfg *Server) {
		cfg.addr = addr
	})
}

func WithServerOption(gsos ...grpc.ServerOption) opts.Option[Server] {
	return opts.OptionFunc[Server](func(cfg *Server) {
		cfg.srvOpts = gsos
	})
}
