package fetchers

import (
	"github.com/transerver/app/configs"
	"github.com/transerver/app/logger"
)

type Redis struct{}

func (f *Redis) Fetch() (configs.Redis, error) {
	return configs.Redis{}, nil
}

func init() {
	logger.Info("redis init fethcer...")
	configs.RegisterCachedFetcher[configs.Redis](&Redis{})
}
