package data

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/transerver/accounts/internal/conf"
	"github.com/transerver/accounts/internal/ent"
	"github.com/transerver/commons/configs"
	"github.com/transerver/commons/logger"
	"github.com/transerver/commons/rs"
	"github.com/xo/dburl"
	"strings"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewAccountRepo,
	NewRsaRepo,
	NewRegionRepo,
)

type Data struct {
	ent   *ent.Client
	redis *rs.Client
	err   error
}

func NewData() (data *Data, cleanup func(), err error) {
	data = &Data{}
	data.connectDatabase()
	data.connectRedis()

	if data.err != nil {
		return nil, nil, data.err
	}

	return data, func() {
		_ = data.ent.Close()
		logger.Sugar().Infof("closing database connection.")
		_ = data.redis.Close()
		logger.Sugar().Infof("closing redis connection.")
	}, data.err
}

func (d *Data) connectRedis() {
	if d.err != nil {
		return
	}

	d.redis, d.err = conf.Bootstrap.Redis.Connect(rs.Config{})
	if d.err != nil {
		d.err = errors.Wrap(d.err, "connect redis error")
	} else {
		logger.Sugar().Infof("[%s] connnect successful",
			strings.Join(conf.Bootstrap.Redis.Address, ", "))
	}
}

func (d *Data) connectDatabase() {
	var drv *sql.Driver
	var url *dburl.URL
	drv, url, d.err = conf.Bootstrap.Database.Connect()
	if d.err != nil {
		d.err = errors.Wrap(d.err, "connect database error")
		return
	}

	opts := []ent.Option{
		ent.Driver(drv),
		ent.Log(func(a ...any) {
			logger.Sugar().Debug(a...)
		}),
	}
	if conf.Bootstrap.Env() == configs.DEV {
		opts = append(opts, ent.Debug())
	}
	d.ent = ent.NewClient(opts...)
	logger.Sugar().Infof("[%s] connect successfully!!!", url.URL.Redacted())
	return
}
