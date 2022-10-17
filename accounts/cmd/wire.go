//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/transerver/accounts/internal/biz"
	"github.com/transerver/accounts/internal/data"
	"github.com/transerver/accounts/internal/service"
	"github.com/transerver/commons/gs"
)

func wireApp() (*gs.Server, func(), error) {
	wire.Build(
		gs.NewGRPCServer,
		NewGRPCOpts,
		biz.ProviderSet,
		data.ProviderSet,
		service.ProviderSet,
	)
	return nil, nil, nil
}

func NewGRPCOpts() []gs.Option {
	return nil
}
