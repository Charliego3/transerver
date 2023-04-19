package configs

type Config struct {
	Target    string
	Address   string
	AccessKey string
	SecretKey string
}

type RedisConfig struct {
}

type Loader[T any] interface {
	FetchConfig() T
}

type Fetcher interface {
	EtcdFetcher
	RedisFetcher
}

type EtcdFetcher interface {
	FetchEtcdConfig()
}

type RedisFetcher interface {
	FetchRedisConfig()
}

type EtcdLoader struct {
}

func (l *EtcdLoader) FetchConfig() Config {
	return Config{}
}

type RedisLoader struct {
}

func (l *RedisLoader) FetchConfig() RedisConfig {
	return RedisConfig{}
}
