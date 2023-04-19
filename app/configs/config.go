package configs

var (
	defaultLoader Loader = nil
)

func SetLoader(loader Loader) {
	defaultLoader = loader
}

func ConfigCenterProps() Config {
	return Config{}
}
