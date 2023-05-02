package fetchers

import "github.com/transerver/app/configs"

type RedisFetcher struct{}

func (f *RedisFetcher) Fetch() (configs.Redis, error) {
	return configs.Redis{}, nil
}
