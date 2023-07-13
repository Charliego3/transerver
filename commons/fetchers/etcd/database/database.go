package fetchers

import (
	"github.com/goccy/go-json"
	"github.com/transerver/mapp/configs"
	"github.com/transerver/mapp/etcdx"
)

type Database struct{}

func (f *Database) Fetch() (cfg configs.Database, err error) {
	resp, e := etcdx.Fetch("")
	if e != nil {
		return cfg, e
	}

	err = json.Unmarshal(resp.Kvs[0].Value, &cfg)
	return
}

func init() {
	configs.RegisterCachedFetcher[configs.Database](&Database{})
}
