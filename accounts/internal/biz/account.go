package biz

import (
	"context"
	"github.com/transerver/accounts/internal/ent"
	"go.uber.org/zap"
)

type AccountRepo interface {
	FindById(ctx context.Context, id int64) (*ent.Account, error)
}

type AccountUsecase struct {
	repo   AccountRepo
	logger *zap.Logger
}

type AccountHelper struct {
	repo   AccountRepo
	logger *zap.Logger
	err    error
}

func NewAccountUsecase(repo AccountRepo, logger *zap.Logger) *AccountUsecase {
	return &AccountUsecase{repo: repo, logger: logger}
}

func (g *AccountUsecase) Helper() *AccountHelper {
	return &AccountHelper{repo: g.repo, logger: g.logger}
}

func (h *AccountHelper) Err() error {
	return h.err
}
