package fetchers

import (
	"errors"
	"os"
	"strings"

	"github.com/charliego93/argsx"
	"github.com/goccy/go-json"
	"github.com/gookit/goutil/strutil"
	"github.com/transerver/app/configs"
	"github.com/transerver/utils"
)

type Etcd struct{}

func (f *Etcd) getByPrefix(key string) string {
	return os.Getenv("ETCD_" + key)
}

func (f *Etcd) Fetch() (configs.Etcd, error) {
	cfgJson := f.getByPrefix("JSON_CONFIG")
	if strutil.IsNotBlank(cfgJson) && cfgJson[0] == '{' {
		var config configs.Etcd
		err := json.Unmarshal(utils.Bytes(cfgJson), &config)
		return config, err
	}

	endpoints := f.getByPrefix("Endpoints")
	if strutil.IsBlank(endpoints) {
		return configs.Etcd{}, errors.New("can not find etcd config from environment")
	}

	var cfg configs.Etcd
	cfg.Endpoints = strings.Split(endpoints, ",")
	cfg.Username = f.getByPrefix("Username")
	cfg.Password = f.getByPrefix("Password")
	cfg.RootCA = f.getByPrefix("RootCA")
	cfg.PemCert = f.getByPrefix("PemCert")
	cfg.PemKey = f.getByPrefix("PemKey")

	cfg.AutoSyncInterval = argsx.New(f.getByPrefix("AutoSyncInterval")).MustDuration()
	cfg.DialTimeout = argsx.New(f.getByPrefix("DialTimeout")).MustDuration()
	cfg.DialKeepAliveTime = argsx.New(f.getByPrefix("DialKeepAliveTime")).MustDuration()
	cfg.DialKeepAliveTimeout = argsx.New(f.getByPrefix("DialKeepAliveTimeout")).MustDuration()
	cfg.MaxCallSendSize = argsx.New(f.getByPrefix("MaxCallSendSize")).MustInt()
	cfg.MaxCallRecvSize = argsx.New(f.getByPrefix("MaxCallRecvSize")).MustInt()
	cfg.PermWithoutStream = argsx.New(f.getByPrefix("PermWithoutStream")).MustBool(false)
	cfg.RejectOldCluster = argsx.New(f.getByPrefix("RejectOldCluster")).MustBool(false)
	return cfg, nil
}

func init() {
	configs.RegisterCachedFetcher[configs.Etcd](&Etcd{})
}
