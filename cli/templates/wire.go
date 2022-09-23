package main

import (
	"github.com/google/wire"
	"github.com/transerver/commons"
)

func wireApp() (*commons.App, func(), error) {
	wire.Build(commons.NewApp)
	return nil, nil, nil
}
