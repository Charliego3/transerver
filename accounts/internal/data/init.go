package data

import (
	"context"
	"database/sql"
	rv9 "github.com/go-redis/redis/v9"
	db "github.com/transerver/accounts/internal/data/sqlc"
	"github.com/transerver/pkg1/configs"
	"github.com/transerver/pkg1/logger"
	"github.com/transerver/pkg1/rs"
	"github.com/xo/dburl"
)

var (
	redisClient *rs.Client
	conn        *sql.DB
	queries     *db.Queries
)

func init() {
	initDatabase()
	initRedis()
}

func initRedis() {
	var err error
	redisClient, err = configs.Bootstrap.Root().Redis.Connect(rs.Config{
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

func initDatabase() {
	var url *dburl.URL
	var err error
	conn, url, err = configs.Bootstrap.Root().Database.Connect()
	if err != nil {
		logger.Sugar().Fatal("connect database error", err)
	}
	logger.Sugar().Infof("connect database: %s", url.Redacted())
	queries = db.New(conn)
}

func Tx(opts ...*sql.TxOptions) (*db.Queries, error) {
	var opt *sql.TxOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	tx, err := conn.BeginTx(context.Background(), opt)
	if err != nil {
		return nil, err
	}
	return queries.WithTx(tx), nil
}

func MustTx(opts ...*sql.TxOptions) *db.Queries {
	tx, err := Tx(opts...)
	if err != nil {
		logger.Sugar().Fatal(err)
	}
	return tx
}
