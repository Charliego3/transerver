package service

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Service interface {
	RegisterGRPC(*grpc.Server)
	RegisterHTTP(*runtime.ServeMux) error

	// Routers returns private route path [grpc, http]
	Routers() ([]string, []string)
}
