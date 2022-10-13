//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/transerver/commons"
	"github.com/transerver/commons/configs"
	"github.com/transerver/commons/gw"
	"github.com/transerver/commons/logger"
)

func wireApp() (*commons.App, func(), error) {
	wire.Build(
		commons.NewApp,
		configs.Parse,
		gw.NewGatewayServer,
		logger.NewLogger,
		providerSet,
	)
	return nil, nil, nil
}
