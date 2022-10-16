package main

import (
	"github.com/Charliego93/go-i18n/v2"
	"github.com/google/wire"
	"github.com/transerver/accounts/internal/biz"
	"github.com/transerver/accounts/internal/conf"
	"github.com/transerver/accounts/internal/data"
	"github.com/transerver/accounts/internal/service"
	"github.com/transerver/commons/configs"
	"github.com/transerver/commons/gs"
	"github.com/transerver/utils"
	"io"
)

var providerSet = wire.NewSet(
	NewCfgOpts,
	NewLoggerWriter,
	biz.ProviderSet,
	conf.ProviderSet,
	data.ProviderSet,
	service.ProviderSet,
)

func main() {
	app, cleanup, err := wireApp()
	if err != nil {
		panic(err)
	}

	defer cleanup()
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func NewCfgOpts() []configs.Option {
	return []configs.Option{
		configs.WithPath("internal/conf/config.yaml"),
	}
}

func NewLoggerWriter() io.Writer {
	return io.Discard
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
