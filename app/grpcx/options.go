package grpcx

import "google.golang.org/grpc"

type Option func(server *GrpcServer)

func WithAddr(addr string) Option {
	return func(g *GrpcServer) {
		g.addr = addr
	}
}

func WithServerOption(opts ...grpc.ServerOption) Option {
	return func(g *GrpcServer) {
		g.srvOpts = append(g.srvOpts, opts...)
	}
}
