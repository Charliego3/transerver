package main

import (
	"github.com/transerver/app/etcdx"
	"github.com/transerver/app/logger"
	_ "github.com/transerver/commons/fetchers/env/etcd"
	_ "github.com/transerver/commons/fetchers/etcd/database"
	_ "github.com/transerver/commons/fetchers/etcd/redis"
)

func main() {
	logger.SetLevel(logger.DebugLevel)
	_, err := etcdx.Fetch("app_config")
	if err != nil {
		logger.Fatal("cant fetch", "err", err)
	}
}
