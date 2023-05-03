package fetchers

import (
	"github.com/goccy/go-json"
	"github.com/transerver/app/configs"
	"github.com/transerver/app/etcdx"
)

type Redis struct{}

func (f *Redis) Fetch() (cfg configs.Redis, err error) {
	resp, e := etcdx.Fetch("")
	if e != nil {
		return cfg, e
	}

	err = json.Unmarshal(resp.Kvs[0].Value, &cfg)
	return
}

func init() {
	configs.RegisterCachedFetcher[configs.Redis](&Redis{})
}
