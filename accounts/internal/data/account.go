package data

import (
	"context"
	"github.com/transerver/accounts/internal/biz"
	"github.com/transerver/accounts/internal/ent"
)

var _ biz.AccountRepo = (*accountRepo)(nil)

type accountRepo struct {
	data *Data
}

func NewAccountRepo(data *Data) biz.AccountRepo {
	return &accountRepo{data: data}
}

func (g *accountRepo) FindById(ctx context.Context, uid int64) (*ent.Account, error) {
	return nil, nil
}
