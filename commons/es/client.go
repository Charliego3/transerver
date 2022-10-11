package es

import (
	"time"

	ev3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

type Client struct {
	*ev3.Client
}

func NewClientWithoutOpts(logger *zap.Logger) (*Client, error) {
	return NewClient(logger)
}

func NewClient(logger *zap.Logger, opts ...Option) (*Client, error) {
	cfg := &ev3.Config{Logger: logger}
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
