package data

import (
	"context"
	"github.com/transerver/accounts/internal/biz"
	"github.com/transerver/accounts/internal/ent"
	"go.uber.org/zap"
)

var _ biz.RsaRepo = (*rsaRepo)(nil)

type rsaRepo struct {
	data   *Data
	logger *zap.Logger
}

func NewRsaRepo(data *Data, logger *zap.Logger) biz.RsaRepo {
	return &rsaRepo{data: data, logger: logger}
}

func (g *rsaRepo) FindById(ctx context.Context, uid int64) (*ent.Rsa, error) {
	return nil, nil
}
