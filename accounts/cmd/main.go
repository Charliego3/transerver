package main

import (
	"github.com/Charliego93/go-i18n/v2"
	"github.com/transerver/commons/gs"
	"github.com/transerver/commons/logger"
	"github.com/transerver/utils"
)

func main() {
	app, cleanup, err := wireApp()
	if err != nil {
		logger.Sugar().Fatal("create accounts app fail", err)
	}

	defer cleanup()
	if err := app.Run(); err != nil {
		logger.Sugar().Fatal("accounts running error", err)
	}
}

func NewGRPCOpts() []gs.Option {
	return []gs.Option{
		gs.WithUnaryServerInterceptor(gs.UnaryI18nInterceptor),
		gs.WithI18nOpts(
			i18n.WithDefaultLanguage(utils.DefaultLanguage),
			i18n.WithLanguageKey("accept-language"),
			i18n.NewLoaderWithPath("internal/i18n"),
		),
	}
}
