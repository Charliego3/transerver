package configs

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/xo/dburl"
	"time"
)

type Database struct {
	DSN     string `json:"dsn" yaml:"dsn"`
	Options struct {
		MaxOpenConns int           `json:"maxOpenConns,omitempty" yaml:"maxOpenConns,omitempty"`
		MaxIdleConns int           `json:"maxIdleConns,omitempty" yaml:"maxIdleConns,omitempty"`
		MaxIdleTime  time.Duration `json:"maxIdleTime,omitempty" yaml:"maxIdleTime,omitempty"`
		MaxLifetime  time.Duration `json:"maxLifetime,omitempty" yaml:"maxLifetime,omitempty"`
	} `json:"options,omitempty" yaml:"options,omitempty"`
}

func (d *Database) Connect() (drv *sql.Driver, url *dburl.URL, err error) {
	url, err = dburl.Parse(d.DSN)
	if err != nil {
		return
	}

	drv, err = sql.Open(url.Driver, url.DSN)
	if err != nil {
		return
	}

	db := drv.DB()
	db.SetMaxIdleConns(d.Options.MaxIdleConns)
	db.SetMaxOpenConns(d.Options.MaxOpenConns)
	db.SetConnMaxIdleTime(d.Options.MaxIdleTime)
	db.SetConnMaxLifetime(d.Options.MaxLifetime)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return
	}
	return
}
