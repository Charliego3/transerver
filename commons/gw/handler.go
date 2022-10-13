package gw

import "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

type Handler struct {
	Method string
	Path   string
	route  runtime.HandlerFunc
}
