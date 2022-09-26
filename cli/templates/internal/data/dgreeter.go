package data

import (
	"context"

	"github.com/transerver/cli/templates/internal/ent"
	"go.uber.org/zap"
)

type GreeterRepo struct {
	data   *Data
	logger *zap.Logger
}

func NewGreeterRepo(data *Data, logger *zap.Logger) *GreeterRepo {
	return &GreeterRepo{data: data, logger: logger}
}

func (g *GreeterRepo) FindById(ctx context.Context, uid int64) (*ent.Greeter, error) {
	return nil, nil
}

func (g *GreeterRepo) FindList(ctx context.Context, t *ent.Greeter) (*ent.Greeter, error) {
	return nil, nil
}

func (g *GreeterRepo) Save(ctx context.Context, t *ent.Greeter) error {
	return nil
}

func (g *GreeterRepo) Delete(ctx context.Context, uid int64) error {
	return nil
}
