package main

import (
	"github.com/transerver/app"
	_ "github.com/transerver/commons/fetchers/env/etcd"
	_ "github.com/transerver/commons/fetchers/etcd/database"
	_ "github.com/transerver/commons/fetchers/etcd/redis"
)

func main() {
	application := app.NewApp()
	err := application.Run()
	if err != nil {
		panic(err)
	}
}
