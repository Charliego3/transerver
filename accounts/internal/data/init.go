package data

import (
	"context"
	rv9 "github.com/go-redis/redis/v9"
	"github.com/transerver/pkg/configs"
	"github.com/transerver/pkg/logger"
	"github.com/transerver/pkg/rs"
)

var rclient *rs.Client

func init() {
	var err error
	rclient, err = configs.Bootstrap.Root().Redis.Connect(rs.Config{
		OnConnect: func(ctx context.Context, cn *rv9.Conn) error {
			cmd := cn.Ping(ctx)
			if cmd.Err() != nil {
				return cmd.Err()
			}
			logger.Sugar().Infof("connect redis: %s", cn.String())
			return nil
		},
	})
	if err != nil {
		logger.Sugar().Fatal(err)
	}
}
