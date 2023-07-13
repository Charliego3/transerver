package app

import (
	"github.com/goccy/go-json"
	"github.com/transerver/mapp/configx"
	"github.com/transerver/mapp/etcdx"
)

type App struct{}

func (f *App) Fetch() (cfg configx.App, err error) {
	resp, e := etcdx.Fetch("")
	if e != nil {
		return cfg, e
	}

	err = json.Unmarshal(resp.Kvs[0].Value, &cfg)
	return
}

func init() {
	configx.RegisterCachedFetcher[configx.App](&App{})
}
