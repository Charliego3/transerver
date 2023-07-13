package es

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"sync"
	"time"

	"github.com/gookit/goutil/strutil"
	"github.com/transerver/mapp/utils"
	"github.com/transerver/pkg1/configs"
	"github.com/transerver/pkg1/logger"

	ev3 "go.etcd.io/etcd/client/v3"
)

type Client struct {
	*ev3.Client
}

func New(opts ...Option) (*Client, error) {
	cfg := &ev3.Config{Logger: logger.Standard()}
	for _, opt := range opts {
		opt(cfg)
	}

	if len(cfg.Endpoints) == 0 {
		cfg.Endpoints = []string{":2379"}
	}
	if cfg.DialTimeout == 0 {
		cfg.DialTimeout = time.Minute
	}

	cli, err := ev3.New(*cfg)
	return &Client{Client: cli}, err
}

var (
	c *Client
	o sync.Once
)

func C() *Client {
	o.Do(func() {
		e := configs.Bootstrap.Root().Etcd
		var tlsc *tls.Config
		if strutil.IsNotBlank(e.RootCA) {
			if utils.AnyBlank(e.PemKey, e.PemCert) {
				logger.Sugar().Fatalf("the certificate path is incorrect, Key: %q, Cert: %q", e.PemKey, e.PemCert)
			}

			etcdCA, err := os.ReadFile(e.RootCA)
			if err != nil {
				logger.Sugar().Fatal(err)
			}

			pair, err := tls.LoadX509KeyPair(e.PemCert, e.PemKey)
			if err != nil {
				logger.Sugar().Fatal(err)
			}

			rootCertPool := x509.NewCertPool()
			rootCertPool.AppendCertsFromPEM(etcdCA)

			tlsc = &tls.Config{
				RootCAs:      rootCertPool,
				Certificates: []tls.Certificate{pair},
			}
		}

		opts := []Option{
			WithEndpoints(e.Endpoints...),
			WithAutoSyncInterval(e.AutoSyncInterval),
			WithDialTimeout(e.DialTimeout),
			WithDialKeepAliveTime(e.DialKeepAliveTime),
			WithDialKeepAliveTimeout(e.DialKeepAliveTimeout),
			WithMaxCallSendSize(e.MaxCallSendSize),
			WithMaxCallRecvSize(e.MaxCallRecvSize),
			WithUsername(e.Username),
			WithPassword(e.Password),
			WithPermWithoutStream(e.PermWithoutStream),
			WithRejectOldCluster(e.RejectOldCluster),
		}

		if tlsc != nil {
			opts = append(opts, WithTLS(tlsc))
		}

		client, err := New(opts...)
		if err != nil {
			logger.Sugar().Fatal("connect etcd server fail", err)
		}

		c = client
		logger.Sugar().Infof("connect etcd: %s", c.Endpoints())
	})
	return c
}
