package data

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/transerver/accounts/internal/ent"
	"github.com/transerver/commons/configs"
	"github.com/transerver/commons/rs"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewAccountRepo,
	NewRsaRepo,
	NewRegionRepo,
)

type Data struct {
	logger    *zap.SugaredLogger
	bootstrap configs.IConfig
	ent       *ent.Client
	redis     *rs.Client
	err       error
}

func NewData(bootstrap configs.IConfig, logger *zap.Logger) (data *Data, cleanup func(), err error) {
	data = &Data{logger: logger.Sugar(), bootstrap: bootstrap}
	data.connectDatabase()
	data.connectRedis(logger)

	if data.err != nil {
		return nil, nil, data.err
	}

	return data, func() {
		_ = data.ent.Close()
		data.logger.Infof("closing database connection.")
		_ = data.redis.Close()
		data.logger.Infof("closing redis connection.")
	}, data.err
}

func (d *Data) connectRedis(logger *zap.Logger) {
	if d.err != nil {
		return
	}

	d.redis, d.err = rs.ConnectRedis(logger, d.bootstrap, rs.Config{})
	if d.err != nil {
		d.err = errors.Wrap(d.err, "connect redis error")
	}
}

func (d *Data) connectDatabase() {
	var drv *sql.Driver
	drv, d.err = configs.ConnectDatabase(d.logger, d.bootstrap)
	if d.err != nil {
		d.err = errors.Wrap(d.err, "connect database error")
		return
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
	return
}
