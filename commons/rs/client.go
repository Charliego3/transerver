package rs

import (
	"context"
	"crypto/tls"
	"github.com/transerver/commons/configs"
	"go.uber.org/zap"
	"net"
	"strings"
	"time"

	rv9 "github.com/go-redis/redis/v9"
)

type Config struct {
	Dialer    func(ctx context.Context, network, addr string) (net.Conn, error)
	OnConnect func(ctx context.Context, cn *rv9.Conn) error
	TLSConfig *tls.Config
}

type Client struct {
	rv9.UniversalClient
	logger *zap.SugaredLogger
}

// NewClientWithoutOpts returns redis UniversalClient wrapper
// using all default prop
func NewClientWithoutOpts(logger *zap.Logger) (*Client, error) {
	return NewClient(logger)
}

// NewClient returns redis UniversalClient wrapper
// if the options is empty using default address with 127.0.0.1:6379
// if not specified DialTimeout default is a minute
func NewClient(logger *zap.Logger, opts ...Option) (*Client, error) {
	cfg := &rv9.UniversalOptions{}
	for _, opt := range opts {
		opt(cfg)
	}

	if len(cfg.Addrs) == 0 {
		cfg.Addrs = []string{"127.0.0.1:6379"}
	}

	if cfg.DialTimeout == 0 {
		cfg.DialTimeout = time.Minute
	}

	cli := rv9.NewUniversalClient(cfg)

	timeout := cfg.ReadTimeout
	if timeout == 0 {
		timeout = cfg.DialTimeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	sugar := logger.Sugar()
	sugar.Infof("[%s] connnect successful", strings.Join(cfg.Addrs, ", "))
	return &Client{UniversalClient: cli, logger: sugar}, nil
}

func NewClientWithConfig(logger *zap.Logger, bootstrap configs.IConfig, conf Config) (*Client, error) {
	r := bootstrap.Redis()
	return NewClient(
		logger,
		WithAddrs(r.Address...),
		WithDB(r.DB),
		WithUsername(r.Username),
		WithPassword(r.Password),
		WithSentinelUsername(r.SentinelUsername),
		WithSentinelPassword(r.SentinelPassword),
		WithMaxRetries(r.MaxRetries),
		WithMinRetryBackoff(r.MinRetryBackoff),
		WithMaxRetryBackoff(r.MaxRetryBackoff),
		WithDialTimeout(r.DialTimeout),
		WithReadTimeout(r.ReadTimeout),
		WithWriteTimout(r.WriteTimeout),
		WithPoolFIFO(r.PoolFIFO),
		WithPoolSize(r.PoolSize),
		WithPoolTimeout(r.PoolTimeout),
		WithMinIdleConns(r.MinIdleConns),
		WithMaxIdleConns(r.MaxIdleConns),
		WithConnMaxIdleTime(r.ConnMaxIdleTime),
		WithConnMaxLifetime(r.ConnMaxLifetime),
		WithMaxRedirects(r.MaxRedirects),
		WithReadOnly(r.ReadOnly),
		WithRouteByLatency(r.RouteByLatency),
		WithRouteRandomly(r.RouteRandomly),
		WithMasterName(r.MasterName),

		WithDialer(conf.Dialer),
		WithOnConnect(conf.OnConnect),
		WithTLSConfig(conf.TLSConfig),
	)
}

func ConnectRedis(logger *zap.Logger, bootstrap configs.IConfig, conf Config) (*Client, error) {
	client, err := NewClientWithConfig(logger, bootstrap, conf)
	if err != nil {
		return nil, err
	}
	return client, nil
}
