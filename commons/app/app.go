package app

import "github.com/transerver/commons/configs"

var Name string

func init() {
	if configs.Bootstrap.Root().Environment != configs.DEV {
		return
	}

	Name = configs.Bootstrap.Root().Name
}
