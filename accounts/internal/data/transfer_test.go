package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/charliego93/flagx"
	_ "github.com/lib/pq"
	db "github.com/transerver/accounts/internal/data/sqlc"
	"github.com/transerver/pkg/logger"
)

var (
	dbURL = flagx.String("config", "postgres://root:root@127.0.0.1:5432/configuration?sslmode=disable")
	cdb   *sql.DB
)

func init() {
	_ = dbURL
	var err error
	cdb, err = sql.Open("postgres", *dbURL)
	if err != nil {
		logger.Sugar().Fatal("can't open database", *dbURL, err)
	}
	err = cdb.Ping()
	if err != nil {
		logger.Sugar().Fatal("ping database error", *dbURL, err)
	}
	logger.Sugar().Infof("connect database: %s", *dbURL)
}

func TestRegions(t *testing.T) {
	rows, err := cdb.Query("SELECT * FROM configuration.public.country_code")
	if err != nil {
		logger.Sugar().Fatal("can't query country_code", err)
	}

	for rows.Next() {
		param := db.RegionCreateParams{Name: json.RawMessage{}}
		err := rows.Scan(&param.Code, &param.Area, &param.Img, &param.Name)
		if err != nil {
			logger.Sugar().Fatal("scan country_code error", err)
		}
		_, err = queries.RegionCreate(context.Background(), param)
		if err != nil {
			logger.Sugar().Fatal("create region error", err)
		}
	}

	_ = rows.Close()
}
