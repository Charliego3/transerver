package config

import (
	"fmt"

	cv2 "github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"github.com/gookit/goutil/strutil"
)

type Config struct {
	paths  []string
	source []byte
	err    error // error with load conetnt, fail fast
}

func (c *Config) Load(source any) {
	if c.err != nil {
		return
	}

	switch v := source.(type) {
	case []byte:
		if len(v) == 0 {
			return
		}
		c.err = cv2.LoadSources(cv2.Yaml, v)
	case []string:
		if len(v) == 0 {
			return
		}
		for _, path := range v {
			c.Load(path)
		}
	case string:
		if strutil.IsBlank(v) {
			return
		}
		c.err = cv2.LoadFiles(v)
	}
}

func ParseWithoutOpts(bootstrap any) (Bootstrap, error) {
	return Parse(bootstrap)
}

// Parse bind cofnig source to Bootstrap instance
// bootstrap must be implement Bootstrap interface
// 1: load from etcd
// 2: load from source option
// 3: load from yaml file
func Parse(bootstrap any, opts ...Option) (Bootstrap, error) {
	if bootstrap == nil {
		return nil, fmt.Errorf("config.Bootstrap target is nil")
	}

	cfg := &Config{}
	for _, opt := range opts {
		opt(cfg)
	}

	cv2.AddDriver(yamlv3.Driver)
	cfg.Load([]byte("")) // from etcd
	cfg.Load(cfg.source)
	cfg.Load(cfg.paths)
	return bootstrap.(Bootstrap), cfg.err
}
