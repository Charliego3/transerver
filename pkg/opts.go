package pkg

import "google.golang.org/grpc"

type Option interface {
	apply()
}

func WithGrpcServer(fn func(server *grpc.Server)) Option {
	return nil
}