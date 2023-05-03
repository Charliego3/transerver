package main

import (
	"github.com/transerver/app/etcdx"
	_ "github.com/transerver/commons/fetchers/env/etcd"
	_ "github.com/transerver/commons/fetchers/etcd/database"
	_ "github.com/transerver/commons/fetchers/etcd/redis"
)

func main() {
	etcdx.C()
}
