package app

import (
	"github.com/goccy/go-json"
	configs "github.com/transerver/mapp/configx"
	"github.com/transerver/mapp/etcdx"
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
