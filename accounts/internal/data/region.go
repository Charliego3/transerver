package data

import (
	"context"
	"github.com/transerver/accounts/internal/biz"
	db "github.com/transerver/accounts/internal/data/sqlc"
)

type regionRepo struct{}

func NewRegionRepo() biz.RegionRepo {
	return &regionRepo{}
}

func (r *regionRepo) List(ctx context.Context) ([]db.Region, error) {
	return queries.RegionList(ctx)
}

func (r *regionRepo) ByCode(ctx context.Context, code string) (db.Region, error) {
	return queries.RegionByCode(ctx, code)
}
