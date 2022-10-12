package configs

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/xo/dburl"
	"go.uber.org/zap"
	"time"
)

type DBConfig struct {
	DSN     string `json:"dsn" yaml:"dsn"`
	Options struct {
		MaxOpenConns int           `json:"maxOpenConns,omitempty" yaml:"maxOpenConns,omitempty"`
		MaxIdleConns int           `json:"maxIdleConns,omitempty" yaml:"maxIdleConns,omitempty"`
		MaxIdleTime  time.Duration `json:"maxIdleTime,omitempty" yaml:"maxIdleTime,omitempty"`
		MaxLifetime  time.Duration `json:"maxLifetime,omitempty" yaml:"maxLifetime,omitempty"`
	} `json:"options,omitempty" yaml:"options,omitempty"`
}

func ConnectDatabase(logger *zap.SugaredLogger, bootstrap IConfig) (drv *sql.Driver, err error) {
	dbc := bootstrap.DB()
	var url *dburl.URL
	url, err = dburl.Parse(dbc.DSN)
	if err != nil {
		return
	}

	drv, err = sql.Open(url.Driver, url.DSN)
	if err != nil {
		return
	}

	db := drv.DB()
	db.SetMaxIdleConns(dbc.Options.MaxIdleConns)
	db.SetMaxOpenConns(dbc.Options.MaxOpenConns)
	db.SetConnMaxIdleTime(dbc.Options.MaxIdleTime)
	db.SetConnMaxLifetime(dbc.Options.MaxLifetime)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return
	}

	logger.Infof("[%s] connect successfully!!!", url.URL.Redacted())
	return drv, nil
}
