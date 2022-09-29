package data

import (
	"context"
	"github.com/transerver/accounts/internal/biz"
	"github.com/transerver/accounts/internal/ent"
	"go.uber.org/zap"
)

var _ biz.AccountRepo = (*accountRepo)(nil)

type accountRepo struct {
	data   *Data
	logger *zap.Logger
}

func NewAccountRepo(data *Data, logger *zap.Logger) biz.AccountRepo {
	return &accountRepo{data: data, logger: logger}
}

func (g *accountRepo) FindById(ctx context.Context, uid int64) (*ent.Account, error) {
	return nil, nil
}
