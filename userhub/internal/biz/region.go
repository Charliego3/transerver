package biz

import (
	"context"
	db "github.com/transerver/accounts/internal/data/sqlc"
)

type RegionRepo interface {
	List(context.Context) ([]db.Region, error)
	ByCode(context.Context, string) (db.Region, error)
}

type RegionUsecase struct {
	repo RegionRepo
}

func NewRegionUsecase(repo RegionRepo) *RegionUsecase {
	return &RegionUsecase{repo: repo}
}

func (g *RegionUsecase) Regions(ctx context.Context) ([]db.Region, error) {
	return g.repo.List(ctx)
}
