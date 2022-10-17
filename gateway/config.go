package main

import (
	"github.com/transerver/commons/configs"
)

// Bootstrap instance
var Bootstrap = &bootstrap{}

type bootstrap struct {
	configs.Base `json:",inline" yaml:",inline"`
}

func init() {
	configs.RegisterBootstrap(Bootstrap)
	configs.Parse(configs.NewFileLoader("config.yaml"))
}
