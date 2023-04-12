package main

import (
	"github.com/Charliego93/go-i18n/v2"
	"github.com/transerver/accounts/internal/conf"
	"github.com/transerver/pkg1/gs"
	"github.com/transerver/pkg1/logger"
)

func main() {
	app, cleanup, err := wireApp()
	if err != nil {
		logger.Sugar().Fatalf("create accounts app fail: %v", err)
	}

	defer cleanup()
	if err := app.Run(); err != nil {
		logger.Sugar().Fatalf("accounts running error: %v", err)
	}
}

func NewGRPCOpts() []gs.Option {
	return []gs.Option{
		gs.WithI18nOpts(i18n.NewLoaderWithPath(conf.Bootstrap.I18nPath)),
	}
}
