package biz

import (
	"context"
	"github.com/transerver/accounts/internal/ent"
)

type RegionRepo interface {
	FindByCode(ctx context.Context, code string) (*ent.Region, error)
	All(ctx context.Context, lang string) (ent.Regions, error)
}

type RegionUsecase struct {
	repo RegionRepo
}

type RegionHelper struct {
	repo RegionRepo
	err  error
}

func NewRegionUsecase(repo RegionRepo) *RegionUsecase {
	return &RegionUsecase{repo: repo}
}

func (g *RegionUsecase) Helper() *RegionHelper {
	return &RegionHelper{repo: g.repo}
}

func (h *RegionHelper) Err() error {
	return h.err
}

func (g *RegionUsecase) Regions(ctx context.Context) (ent.Regions, error) {
	return g.repo.All(ctx, "")
}
