//go:build wireinject
// +build wireinject

//go:generate go run github.com/google/wire/cmd/wire
package main

import (
	"github.com/google/wire"
	"github.com/transerver/accounts/internal/biz"
	"github.com/transerver/accounts/internal/data"
	"github.com/transerver/accounts/internal/service"
	"github.com/transerver/pkg/gs"
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
