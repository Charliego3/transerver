package main

import (
	"github.com/transerver/pkg1/configs"
	"github.com/transerver/pkg1/resolver"
)

var (
	// check type
	_ configs.IBootstrap = (*bootstrap)(nil)

	// Bootstrap instance
	Bootstrap = &bootstrap{}
)

type bootstrap struct {
	configs.Base `json:",inline" yaml:",inline"`
	Services     []struct {
		Target string                 `json:"target" yaml:"target"`
		Config resolver.ServiceConfig `json:"config" yaml:"config"`
	} `json:"services" yaml:"services"`
}

func init() {
	configs.ParseConfig(
		Bootstrap,
		configs.NewFileLoader("config.yaml"),
	)
}
