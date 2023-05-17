package main

import (
	"github.com/charliego93/logger"
	"github.com/transerver/app/etcdx"
	_ "github.com/transerver/commons/fetchers/env/etcd"
	_ "github.com/transerver/commons/fetchers/etcd/database"
	_ "github.com/transerver/commons/fetchers/etcd/redis"
)

func main() {
	logger.SetLevel(logger.LevelDebug)
	_, err := etcdx.Fetch("app_config")
	if err != nil {
		logger.Fatal("cant fetch", "err", err)
	}
}
