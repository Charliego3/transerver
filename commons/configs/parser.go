package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type Config struct {
	paths  []string
	source []byte
	reader io.Reader
	err    error // error with load content, fail fast
}

func (c *Config) unmarshal(buf []byte, target any) {
	if c.err != nil || len(buf) == 0 {
		return
	}

	c.err = yaml.Unmarshal(buf, target)
}

func (c *Config) readf(pidx int) ([]byte, bool) {
	if c.err != nil || len(c.paths)-1 > pidx {
		return nil, false
	}

	var buf []byte
	buf, c.err = os.ReadFile(c.paths[pidx])
	if c.err != nil {
		return nil, false
	}
	return buf, true
}

func (c *Config) readr() []byte {
	if c.err != nil || c.reader == nil {
		return nil
	}

	var buf []byte
	for {
		tmp := make([]byte, 1024)
		_, err := c.reader.Read(tmp)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil
		}
		buf = append(buf, tmp...)
	}
	return buf
}

func ParseWithoutOpts(bootstrap any) (IConfig, error) {
	return Parse(bootstrap)
}

// Parse bind config source to IConfig instance
// bootstrap must be implemented IConfig interface
func Parse(bootstrap any, opts ...Option) (IConfig, error) {
	if bootstrap == nil {
		return nil, fmt.Errorf("config.IConfig target is nil")
	}

	if _, ok := bootstrap.(IConfig); !ok {
		return nil, fmt.Errorf("bootstrap does not implemented configs.IConfig")
	}

	cfg := &Config{}
	for _, opt := range opts {
		opt(cfg)
	}

	cfg.unmarshal(cfg.source, bootstrap)
	cfg.unmarshal(cfg.readr(), bootstrap)
	for i := range cfg.paths {
		if buf, ok := cfg.readf(i); ok {
			cfg.unmarshal(buf, bootstrap)
		}
	}
	return bootstrap.(IConfig), cfg.err
}
