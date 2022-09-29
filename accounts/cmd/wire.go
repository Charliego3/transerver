//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/transerver/commons"
	"github.com/transerver/commons/configs"
	"github.com/transerver/commons/gs"
	"github.com/transerver/commons/hs"
	"github.com/transerver/commons/logger"
)

func wireApp() (*commons.App, func(), error) {
	wire.Build(
		commons.NewApp,
		configs.Parse,
		gs.NewGRPCServer,
		hs.NewHTTPServerWithoutHandlers,
		logger.NewLogger,
		providerSet,
	)
	return nil, nil, nil
}
