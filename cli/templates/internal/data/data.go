package data

import (
	"github.com/google/wire"
	"github.com/transerver/commons/configs"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewGreeterRepo,
)

type Data struct {
	logger *zap.Logger
}

func NewData(bootstrap configs.Bootstrap, logger *zap.Logger) *Data {
	return &Data{logger: logger}
}
