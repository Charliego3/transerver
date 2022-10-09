package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	rv9 "github.com/go-redis/redis/v9"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"github.com/transerver/accounts/internal/ent"
	"github.com/transerver/commons/configs"
	"github.com/transerver/commons/rs"
	"github.com/xo/dburl"
	"go.uber.org/zap"
	"time"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewAccountRepo,
	NewRsaRepo,
)

type Data struct {
	logger    *zap.SugaredLogger
	bootstrap configs.IConfig
	ent       *ent.Client
	redis     *rs.Client
	err       error
}

func NewData(bootstrap configs.IConfig, logger *zap.Logger) (*Data, func(), error) {
	data := &Data{logger: logger.Sugar(), bootstrap: bootstrap}
	cleanDB := data.connectDatabase()
	cleanRedis := data.connectRedis()
	return data, func() {
		if cleanDB != nil {
			_ = cleanDB()
			data.logger.Infof("closing database connection.")
		}
		if cleanRedis != nil {
			_ = cleanRedis()
			data.logger.Infof("closing redis connection.")
		}
	}, data.err
}

func (d *Data) connectRedis() func() error {
	if d.err != nil {
		return nil
	}

	client, err := rs.NewClientWithConfig(d.bootstrap, rs.Config{
		OnConnect: func(ctx context.Context, cn *rv9.Conn) error {
			d.logger.Infof("[%s] connnect successful", cn.String())
			return nil
		},
	})
	if err != nil {
		return nil
	}
	d.redis = client
	return client.Close
}

func (d *Data) connectDatabase() func() error {
	dbc := d.bootstrap.DB()
	var url *dburl.URL
	url, d.err = dburl.Parse(dbc.DSN)
	if d.err != nil {
		return nil
	}

	var drv *sql.Driver
	drv, d.err = sql.Open(url.Driver, url.DSN)
	if d.err != nil {
		return nil
	}

	db := drv.DB()
	db.SetMaxIdleConns(dbc.Options.MaxIdleConns)
	db.SetMaxOpenConns(dbc.Options.MaxOpenConns)
	db.SetConnMaxIdleTime(dbc.Options.MaxIdleTime)
	db.SetConnMaxLifetime(dbc.Options.MaxLifetime)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if d.err = db.PingContext(ctx); d.err != nil {
		return nil
	}

	opts := []ent.Option{
		ent.Driver(drv),
		ent.Log(func(a ...any) {
			d.logger.Debug(a...)
		}),
	}
	if d.bootstrap.Env() == configs.DEV {
		opts = append(opts, ent.Debug())
	}
	d.ent = ent.NewClient(opts...)
	d.logger.Infof("[%s] connect successfully!!!", url.URL.Redacted())
	return d.ent.Close
}
