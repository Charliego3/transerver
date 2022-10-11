package main

import (
	"errors"
	"github.com/google/wire"
	"github.com/transerver/accounts/internal/biz"
	"github.com/transerver/accounts/internal/conf"
	"github.com/transerver/accounts/internal/data"
	"github.com/transerver/accounts/internal/service"
	"github.com/transerver/commons/configs"
	"github.com/transerver/commons/hs"
	"io"
	"net/http"
)

var providerSet = wire.NewSet(
	NewCfgOpts,
	NewLoggerWriter,
	NewHTTPOptions,
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
	app.Run()
}

func NewCfgOpts() []configs.Option {
	return []configs.Option{
		configs.WithPath("internal/conf/config.yaml"),
	}
}

func NewHTTPOptions() []hs.Option {
	return []hs.Option{
		hs.WithAuthFunc(func(r *http.Request) error {
			return errors.New("auth not pass")
		}),
	}
}

func NewLoggerWriter() io.Writer {
	return io.Discard
}
