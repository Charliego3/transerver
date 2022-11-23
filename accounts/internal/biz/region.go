package biz

import (
	"context"
	db "github.com/transerver/accounts/internal/db/sqlc"
)

type RegionUsecase struct{}

func NewRegionUsecase() *RegionUsecase {
	return &RegionUsecase{}
}

func (g *RegionUsecase) Regions(ctx context.Context) ([]db.Region, error) {
	return db.Query().RegionList(ctx)
}
