package db

import (
	"context"
	"database/sql"
	"github.com/transerver/pkg/configs"
	"github.com/transerver/pkg/logger"
)

var queries *Queries

func init() {
	db, url, err := configs.Bootstrap.Root().Database.Connect()
	if err != nil {
		logger.Sugar().Fatal("connect database error", err)
	}
	logger.Sugar().Infof("connect database: %s", url.Redacted())
	queries = New(db)
}

func Query() *Queries {
	return queries
}

func Tx(opts *sql.TxOptions) (*Queries, error) {
	db := queries.db.(*sql.DB)
	tx, err := db.BeginTx(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	return queries.WithTx(tx), nil
}

func MustTx(opt *sql.TxOptions) *Queries {
	tx, err := Tx(opt)
	if err != nil {
		logger.Sugar().Fatal(err)
	}
	return tx
}
