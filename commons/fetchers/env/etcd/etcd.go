package fetchers

import (
	"os"
	"strings"

	"github.com/charliego93/argsx"
	"github.com/goccy/go-json"
	"github.com/gookit/goutil/strutil"
	"github.com/transerver/mapp/configx"
	"github.com/transerver/mapp/utils"
)

type Etcd struct{}

func (f *Etcd) getByPrefix(key string) string {
	return os.Getenv("ETCD_" + key)
}

func (f *Etcd) Fetch() (configx.Etcd, error) {
	cfgJson := f.getByPrefix("JSON_CONFIG")
	var cfg configx.Etcd
	if strutil.IsNotBlank(cfgJson) && cfgJson[0] == '{' {
		err := json.Unmarshal(utils.Bytes(cfgJson), &cfg)
		return cfg, err
	}

	cfg.Endpoints = strings.Split(f.getByPrefix("Endpoints"), ",")
	cfg.Username = f.getByPrefix("Username")
	cfg.Password = f.getByPrefix("Password")
	cfg.RootCA = f.getByPrefix("RootCA")
	cfg.PemCert = f.getByPrefix("PemCert")
	cfg.PemKey = f.getByPrefix("PemKey")
	cfg.AutoSyncInterval = argsx.NewV(f.getByPrefix("AutoSyncInterval")).MustDuration()
	cfg.DialTimeout = argsx.NewV(f.getByPrefix("DialTimeout")).MustDuration()
	cfg.DialKeepAliveTime = argsx.NewV(f.getByPrefix("DialKeepAliveTime")).MustDuration()
	cfg.DialKeepAliveTimeout = argsx.NewV(f.getByPrefix("DialKeepAliveTimeout")).MustDuration()
	cfg.MaxCallSendSize = argsx.NewV(f.getByPrefix("MaxCallSendSize")).MustInt()
	cfg.MaxCallRecvSize = argsx.NewV(f.getByPrefix("MaxCallRecvSize")).MustInt()
	cfg.PermWithoutStream = argsx.NewV(f.getByPrefix("PermWithoutStream")).MustBool(false)
	cfg.RejectOldCluster = argsx.NewV(f.getByPrefix("RejectOldCluster")).MustBool(false)
	return cfg, nil
}

func init() {
	configx.RegisterCachedFetcher[configx.Etcd](&Etcd{})
}
