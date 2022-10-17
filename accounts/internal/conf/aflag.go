package conf

import (
	"flag"
)

var (
	path *string
)

func init() {
	path = flag.String("config", "internal/conf/config.yaml", "config path")
	flag.Parse()
}
