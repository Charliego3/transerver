package app

import (
	"github.com/goccy/go-json"
	"github.com/transerver/app/configs"
	"github.com/transerver/app/etcdx"
)

type App struct{}

func (f *App) Fetch() (cfg configs.App, err error) {
	resp, e := etcdx.Fetch("")
	if e != nil {
		return cfg, e
	}

	err = json.Unmarshal(resp.Kvs[0].Value, &cfg)
	return
}

func init() {
	configs.RegisterCachedFetcher[configs.App](&App{})
}
