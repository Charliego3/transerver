package gs

import (
	"google.golang.org/grpc"
)

type Service interface {
	RegisterGRPC(*grpc.Server)
	// RegisterHTTP(*runtime.ServeMux) error

	// Routers returns private route path [grpc, http]
	// Routers() ([]string, []string)
}
