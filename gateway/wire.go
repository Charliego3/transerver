//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/transerver/commons/configs"
	"github.com/transerver/commons/gw"
	"github.com/transerver/commons/logger"
)

func wireApp() (*gw.Server, func(), error) {
	wire.Build(
		configs.Parse,
		gw.NewGatewayServer,
		logger.NewLogger,
		providerSet,
	)
	return nil, nil, nil
}
