package etcdx

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"sync"
	"time"

	"github.com/gookit/goutil/strutil"
	"github.com/transerver/app/configs"
	"github.com/transerver/app/logger"
	"github.com/transerver/utils"

	ev3 "go.etcd.io/etcd/client/v3"
)

type Client struct {
	*ev3.Client
}

// New returns etcd client from options
func New(opts ...Option) (*Client, error) {
	// TODO: etcd config logger
	cfg := &ev3.Config{}
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

// C returns etcd client from config
func C() *Client {
	o.Do(func() {
		cfg, err := configs.Fetch[configs.Etcd]()
		if err != nil {
			logger.Fatal("not resolved etcd config", "err", err)
		}

		var tlsConfig *tls.Config
		if strutil.IsNotBlank(cfg.RootCA) {
			if utils.AnyBlank(cfg.PemKey, cfg.PemCert) {
				logger.Fatalf("the certificate path is incorrect, Key: %q, Cert: %q", cfg.PemKey, cfg.PemCert)
			}

			etcdCA, err := os.ReadFile(cfg.RootCA)
			if err != nil {
				logger.Fatal("read etcd RootCA fail", "err", err)
			}

			pair, err := tls.LoadX509KeyPair(cfg.PemCert, cfg.PemKey)
			if err != nil {
				logger.Fatal("load etcd CA cert and key fail", "err", err)
			}

			rootCertPool := x509.NewCertPool()
			rootCertPool.AppendCertsFromPEM(etcdCA)

			tlsConfig = &tls.Config{
				RootCAs:      rootCertPool,
				Certificates: []tls.Certificate{pair},
			}
		}

		opts := []Option{
			WithEndpoints(cfg.Endpoints...),
			WithAutoSyncInterval(cfg.AutoSyncInterval),
			WithDialTimeout(cfg.DialTimeout),
			WithDialKeepAliveTime(cfg.DialKeepAliveTime),
			WithDialKeepAliveTimeout(cfg.DialKeepAliveTimeout),
			WithMaxCallSendSize(cfg.MaxCallSendSize),
			WithMaxCallRecvSize(cfg.MaxCallRecvSize),
			WithUsername(cfg.Username),
			WithPassword(cfg.Password),
			WithPermWithoutStream(cfg.PermWithoutStream),
			WithRejectOldCluster(cfg.RejectOldCluster),
		}

		if tlsConfig != nil {
			opts = append(opts, WithTLS(tlsConfig))
		}

		client, err := New(opts...)
		if err != nil {
			logger.Fatal("connect etcd server fail", "err", err)
		}

		c = client
		logger.Infof("connect etcd: %s", c.Endpoints())
	})
	return c
}
