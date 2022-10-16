//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/transerver/commons/configs"
	"github.com/transerver/commons/gs"
	"github.com/transerver/commons/logger"
)

func wireApp() (*gs.Server, func(), error) {
	wire.Build(
		configs.Parse,
		gs.NewGRPCServer,
		NewGRPCOpts,
		logger.NewLogger,
		providerSet,
	)
	return nil, nil, nil
}
