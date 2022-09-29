package main

import (
	"github.com/google/wire"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/accounts/internal/biz"
	"github.com/transerver/accounts/internal/conf"
	"github.com/transerver/accounts/internal/data"
	"github.com/transerver/accounts/internal/service"
	"github.com/transerver/commons/configs"
	"github.com/transerver/commons/hs"
	"google.golang.org/grpc"
	"io"
)

var providerSet = wire.NewSet(
	NewCfgOpts,
	NewGRPCOpts,
	NewHTTPServeMuxOpts,
	NewLoggerWriter,
	biz.ProviderSet,
	conf.ProviderSet,
	data.ProviderSet,
	service.ProviderSet,
)

func main() {
	app, cleanup, err := wireApp()
	if err != nil {
		panic(err)
	}

	defer cleanup()
	app.Run()
}

func NewCfgOpts() []configs.Option {
	return []configs.Option{
		configs.WithPath("internal/conf/config.yaml"),
	}
}

func NewGRPCOpts() []grpc.ServerOption {
	return []grpc.ServerOption{}
}

func NewHTTPServeMuxOpts() []runtime.ServeMuxOption {
	return hs.DefaultServeMuxOpts()
}

func NewLoggerWriter() io.Writer {
	return io.Discard
}
