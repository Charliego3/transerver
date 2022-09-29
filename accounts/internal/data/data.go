package data

import (
	"github.com/google/wire"
	"github.com/transerver/commons/configs"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewAccountRepo,
	NewRsaRepo,
)

type Data struct {
	logger *zap.Logger
}

func NewData(bootstrap configs.IConfig, logger *zap.Logger) (*Data, func(), error) {
	return &Data{logger: logger}, func() {}, nil
}
