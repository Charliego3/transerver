package gw

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type DialerFunc func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

type Dialer struct {
	Target  string
	Opts    []grpc.DialOption
	Servers []DialerFunc
}

func NewDialer(target string, servers []DialerFunc, opts ...grpc.DialOption) Dialer {
	return Dialer{
		Target:  target,
		Servers: servers,
		Opts:    opts,
	}
}
