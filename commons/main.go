package main

import (
	"github.com/charliego93/logger"
	_ "github.com/transerver/commons/fetchers/env/etcd"
	_ "github.com/transerver/commons/fetchers/etcd/database"
	_ "github.com/transerver/commons/fetchers/etcd/redis"
	"github.com/transerver/mapp/etcdx"
)

func main() {
	logger.SetLevel(logger.LevelDebug)
	_, err := etcdx.Fetch("app_config")
	if err != nil {
		logger.Fatal("cant fetch", "err", err)
	}
}
