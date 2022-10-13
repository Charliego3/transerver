package biz

import (
	"context"
	"github.com/transerver/accounts/internal/ent"
	"go.uber.org/zap"
)

type RegionRepo interface {
	FindByCode(ctx context.Context, code string) (*ent.Region, error)
}

type RegionUsecase struct {
	repo   RegionRepo
	logger *zap.Logger
}

type RegionHelper struct {
	repo   RegionRepo
	logger *zap.Logger
	err    error
}

func NewRegionUsecase(repo RegionRepo, logger *zap.Logger) *RegionUsecase {
	return &RegionUsecase{repo: repo, logger: logger}
}

func (g *RegionUsecase) Helper() *RegionHelper {
	return &RegionHelper{repo: g.repo, logger: g.logger}
}

func (h *RegionHelper) Err() error {
	return h.err
}
