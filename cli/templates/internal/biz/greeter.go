package biz

import (
	"github.com/transerver/cli/templates/internal/ent"
	"github.com/transerver/commons"
	"go.uber.org/zap"
)

type GreeterUsecase struct {
	repo   commons.Repository[ent.Greeter]
	logger *zap.Logger
}

type Helper struct {
	repo   commons.Repository[ent.Greeter]
	logger *zap.Logger
	err    error
}

func NewGreeterUsecase(repo commons.Repository[ent.Greeter], logger *zap.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, logger: logger}
}

func (g *GreeterUsecase) Helper() *Helper {
	return &Helper{repo: g.repo, logger: g.logger}
}

func (h *Helper) Err() error {
	return h.err
}
