package configs

import (
	"os"
	"path/filepath"

	"github.com/charliego93/argsx"
	"github.com/goccy/go-json"
	"github.com/gookit/goutil/fsutil"
	"github.com/transerver/app/logger"
	"gopkg.in/yaml.v3"
)

type EmbeddedConfig struct {
	App      *App      `json:"app,omitempty" yaml:"app,omitempty"`
	Etcd     *Etcd     `json:"etcd,omitempty" yaml:"etcd,omitempty"`
	Database *Database `json:"database,omitempty" yaml:"database,omitempty"`
	Redis    *Redis    `json:"redis,omitempty" yaml:"redis,omitempty"`
}

var (
	disabled bool
	instance EmbeddedConfig
)

func init() {
	configPath := argsx.Fetch("config").MustString("./config.yaml")
	disabled = !fsutil.FileExist(configPath)
	if disabled {
		return
	}

	bs, err := os.ReadFile(configPath)
	if err != nil {
		logger.Fatal("read config file", "path", configPath, "err", err)
	}

	switch filepath.Ext(configPath) {
	case ".yaml":
		err = yaml.Unmarshal(bs, &instance)
	case ".json":
		err = json.Unmarshal(bs, &instance)
	}

	if err != nil {
		logger.Fatal("failed to load default config from file", "file", configPath, "err", err)
	}

	register[Etcd](instance.Etcd, &embeddedEtcdFetcher{})
	register[App](instance.App, &embeddedAppFetcher{})
	register[Redis](instance.Redis, &embeddedRedisFetcher{})
	register[Database](instance.Database, &embeddedDatabaseFetcher{})
}

// register register fetcher to fetchers if obj is not nil
func register[T any](obj *T, fetcher Fetcher[T]) {
	if obj == nil {
		return
	}

	RegisterFetcher(fetcher)
}
