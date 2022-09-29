package biz

import (
	"context"
	"github.com/transerver/accounts/internal/ent"
	"go.uber.org/zap"
)

type RsaRepo interface {
	FindById(ctx context.Context, id int64) (*ent.Rsa, error)
}

type RsaUsecase struct {
	repo   RsaRepo
	logger *zap.Logger
}

type RsaHelper struct {
	repo   RsaRepo
	logger *zap.Logger
	err    error
}

func NewRsaUsecase(repo RsaRepo, logger *zap.Logger) *RsaUsecase {
	return &RsaUsecase{repo: repo, logger: logger}
}

func (g *RsaUsecase) Helper() *RsaHelper {
	return &RsaHelper{repo: g.repo, logger: g.logger}
}

func (h *RsaHelper) Err() error {
	return h.err
}
