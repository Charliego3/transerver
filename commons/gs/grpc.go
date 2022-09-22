package gs

import (
	"github.com/transerver/commons"
	"google.golang.org/grpc"
)

func NewGRPCServerWithoutOpts(services []commons.Service) *grpc.Server {
	return NewGRPCServer(services)
}

func NewGRPCServer(services []commons.Service, opts ...grpc.ServerOption) *grpc.Server {
	gs := grpc.NewServer(opts...)

	for _, service := range services {
		service.RegisterGRPC(gs)
	}
	return gs
}
