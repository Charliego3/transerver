package data

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
	"github.com/transerver/accounts/internal/biz"
	"go.uber.org/zap"
	"time"
)

var _ biz.RsaRepo = (*rsaRepo)(nil)

type rsaRepo struct {
	data   *Data
	logger *zap.Logger
}

func NewRsaRepo(data *Data, logger *zap.Logger) biz.RsaRepo {
	return &rsaRepo{data: data, logger: logger}
}

func (g *rsaRepo) Fetch(requestId string) (*biz.RsaObj, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	cmd := g.data.redis.Get(ctx, requestId)
	if cmd.Err() == redis.Nil {
		return nil, nil
	}
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	var rsaObj = &biz.RsaObj{}
	err := cmd.Scan(rsaObj)
	return rsaObj, err
}

func (g *rsaRepo) Store(requestId string, expiration time.Duration, obj *biz.RsaObj) error {
	status, err := g.data.redis.Set(context.Background(), requestId, obj, expiration).Result()
	if err != nil {
		return err
	}
	if "OK" != status {
		return errors.New("store rsa obj error")
	}
	return nil
}
