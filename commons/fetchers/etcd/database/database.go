package fetchers

import (
	"github.com/transerver/app/configs"
	"github.com/transerver/app/logger"
)

type Database struct{}

func (f *Database) Fetch() (configs.Database, error) {
	return configs.Database{}, nil
}

func init() {
	logger.Info("database init fethcer...")
	configs.RegisterCachedFetcher[configs.Database](&Database{})
}
