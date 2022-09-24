package main

import (
	"github.com/google/wire"
	"github.com/transerver/commons"
	"github.com/transerver/commons/configs"
	"github.com/transerver/commons/etcds"
	"github.com/transerver/commons/gs"
	"github.com/transerver/commons/hs"
	"github.com/transerver/commons/logger"
	"github.com/transerver/commons/redis"
)

func wireApp() (*commons.App, func(), error) {
	wire.Build(
		commons.NewApp,
		configs.Parse,
		configs.ParseWithoutOpts,
		etcds.NewClient,
		etcds.NewClientWithoutOpts,
		gs.NewGRPCServer,
		gs.NewGRPCServerWithoutOpts,
		hs.NewHTTPServerWithOptions,
		hs.NewHTTPServerWithoutMuxOpts,
		hs.DefaultOpts,
		hs.NewHTTPServerWithoutOpts,
		logger.NewLoggerWithoutWriter,
		logger.NewLogger,
		redis.NewClientWithoutOpts,
		redis.NewClient,
		providerSet,
	)
	return nil, nil, nil
}
