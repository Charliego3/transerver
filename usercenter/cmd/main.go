package main

import (
	_ "github.com/transerver/commons/fetchers/env/etcd"
	_ "github.com/transerver/commons/fetchers/etcd/database"
	_ "github.com/transerver/commons/fetchers/etcd/redis"
	app "github.com/transerver/mapp"
)

func main() {
	application := app.NewApp(
		app.WithAddr("tcp", ":9001"),
	)
	err := application.Run()
	if err != nil {
		panic(err)
	}
}
