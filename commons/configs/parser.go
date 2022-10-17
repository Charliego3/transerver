package configs

import (
	"gopkg.in/yaml.v3"
)

func Parse(loaders ...Loader) {
	if Bootstrap == nil {
		panic("config.IBootstrap target is nil")
	}

	for _, loader := range loaders {
		buf, err := loader.Load()
		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(buf, Bootstrap)
		if err != nil {
			panic(err)
		}
	}
}
