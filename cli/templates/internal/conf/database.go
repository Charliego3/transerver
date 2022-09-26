package conf

import "time"

type Database struct {
	DSN     string `json:"dsn,omitempty" yaml:"dsn"`
	Options struct {
		MaxOpenConns int           `json:"maxOpenConns,omitempty" yaml:"maxOpenConns"`
		MaxIdleConns int           `json:"maxIdleConns,omitempty" yaml:"maxIdleConns"`
		MaxIdleTime  time.Duration `json:"maxIdleTime,omitempty" yaml:"maxIdleTime"`
		MaxLifetime  time.Duration `json:"maxLifetime,omitempty" yaml:"maxLifetime"`
	} `json:"options" yaml:"options"`
}
