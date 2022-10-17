package configs

import (
	"github.com/transerver/commons/rs"
	"time"
)

type Redis struct {
	// Either a single address or a seed list of host:port addresses
	// of cluster/sentinel nodes.
	Address []string `yaml:"address,omitempty"`

	// Database to be selected after connecting to the server.
	// Only single-node and failover clients.
	DB int `yaml:"db,omitempty"`

	// Common options
	Username         string `yaml:"username,omitempty"`
	Password         string `yaml:"password,omitempty"`
	SentinelUsername string `yaml:"sentinelUsername,omitempty"`
	SentinelPassword string `yaml:"sentinelPassword,omitempty"`

	MaxRetries      int           `yaml:"maxRetries,omitempty"`
	MinRetryBackoff time.Duration `yaml:"minRetryBackoff,omitempty"`
	MaxRetryBackoff time.Duration `yaml:"maxRetryBackoff,omitempty"`

	DialTimeout  time.Duration `yaml:"dialTimeout,omitempty"`
	ReadTimeout  time.Duration `yaml:"readTimeout,omitempty"`
	WriteTimeout time.Duration `yaml:"writeTimeout,omitempty"`

	PoolFIFO    bool          `yaml:"poolFIFO,omitempty"`
	PoolSize    int           `yaml:"poolSize,omitempty"`
	PoolTimeout time.Duration `yaml:"poolTimeout,omitempty"`

	MinIdleConns int `yaml:"minIdleConns,omitempty"`
	MaxIdleConns int `yaml:"MaxIdleConns,omitempty"`

	ConnMaxIdleTime time.Duration `yaml:"connMaxIdleTime,omitempty"`
	ConnMaxLifetime time.Duration `yaml:"ConnMaxLifetime,omitempty"`

	// Only cluster clients
	MaxRedirects   int  `yaml:"maxRedirects,omitempty"`
	ReadOnly       bool `yaml:"readOnly,omitempty"`
	RouteByLatency bool `yaml:"routeByLatency,omitempty"`
	RouteRandomly  bool `yaml:"routeRandomly,omitempty"`

	// The sentinel master name
	// Only failover clients
	MasterName string `yaml:"masterName,omitempty"`
}

func (r *Redis) Connect(opt rs.Config) (*rs.Client, error) {
	return rs.New(
		rs.WithAddrs(r.Address...),
		rs.WithDB(r.DB),
		rs.WithUsername(r.Username),
		rs.WithPassword(r.Password),
		rs.WithSentinelUsername(r.SentinelUsername),
		rs.WithSentinelPassword(r.SentinelPassword),
		rs.WithMaxRetries(r.MaxRetries),
		rs.WithMinRetryBackoff(r.MinRetryBackoff),
		rs.WithMaxRetryBackoff(r.MaxRetryBackoff),
		rs.WithDialTimeout(r.DialTimeout),
		rs.WithReadTimeout(r.ReadTimeout),
		rs.WithWriteTimout(r.WriteTimeout),
		rs.WithPoolFIFO(r.PoolFIFO),
		rs.WithPoolSize(r.PoolSize),
		rs.WithPoolTimeout(r.PoolTimeout),
		rs.WithMinIdleConns(r.MinIdleConns),
		rs.WithMaxIdleConns(r.MaxIdleConns),
		rs.WithConnMaxIdleTime(r.ConnMaxIdleTime),
		rs.WithConnMaxLifetime(r.ConnMaxLifetime),
		rs.WithMaxRedirects(r.MaxRedirects),
		rs.WithReadOnly(r.ReadOnly),
		rs.WithRouteByLatency(r.RouteByLatency),
		rs.WithRouteRandomly(r.RouteRandomly),
		rs.WithMasterName(r.MasterName),

		rs.WithDialer(opt.Dialer),
		rs.WithOnConnect(opt.OnConnect),
		rs.WithTLSConfig(opt.TLSConfig),
	)
}
