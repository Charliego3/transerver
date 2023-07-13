package main

import (
	"context"

	_ "github.com/transerver/commons/fetchers/env/etcd"
	_ "github.com/transerver/commons/fetchers/etcd/database"
	_ "github.com/transerver/commons/fetchers/etcd/redis"
	"github.com/transerver/mapp"
)

func main() {
	application := mapp.NewApp(
		mapp.WithAddr("tcp", ":9001"),
	)
	err := application.Run(context.Background())
	if err != nil {
		panic(err)
	}
}
