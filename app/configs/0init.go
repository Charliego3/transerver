package configs

import (
	"github.com/transerver/app/logger"
	"go.etcd.io/etcd/client/pkg/v3/fileutil"
	"os"
	"path/filepath"
)

var disableDefault bool

func init() {
	logger.Info(filepath.Ext("./config.yaml"))
	disableDefault = !fileutil.Exist("./config.yaml")
	if disableDefault {
		return
	}

	bs, err := os.ReadFile("./config.yaml")
	if err != nil {
		logger.Fatal("read config file", "path", "./config.yaml", "err", err)
	}

	_ = bs
}
