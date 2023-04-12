package configs

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
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

func (d *Database) Connect() (db *sql.DB, url *dburl.URL, err error) {
	url, err = dburl.Parse(d.DSN)
	if err != nil {
		return
	}

	db, err = sql.Open(url.Driver, url.DSN)
	if err != nil {
		return
	}

	db.SetMaxIdleConns(d.Options.MaxIdleConns)
	db.SetMaxOpenConns(d.Options.MaxOpenConns)
	db.SetConnMaxIdleTime(d.Options.MaxIdleTime)
	db.SetConnMaxLifetime(d.Options.MaxLifetime)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	err = db.PingContext(ctx)
	return
}
