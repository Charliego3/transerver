//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/transerver/commons"
	"github.com/transerver/commons/configs"
	"github.com/transerver/commons/gs"
	"github.com/transerver/commons/hs"
	"github.com/transerver/commons/logger"
	"google.golang.org/grpc"
)

func wireApp() (*commons.App, func(), error) {
	wire.Build(
		commons.NewApp,
		configs.Parse,
		gs.NewGRPCServer,
		NewGRPCOpts,
		hs.NewHTTPServer,
		logger.NewLogger,
		providerSet,
	)
	return nil, nil, nil
}

func NewGRPCOpts() []gs.Option {
	return []gs.Option{
		gs.WithAuthFunc(func(ctx context.Context) (context.Context, error) {
			if method, ok := grpc.Method(ctx); ok {
				fmt.Println("GRPC Method:", method)
			}
			return ctx, nil
		}),
	}
}
