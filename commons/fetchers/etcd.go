package fetchers

import "github.com/transerver/app/configs"

type EtcdFetcher struct{}

func (f *EtcdFetcher) Fetch() (configs.Etcd, error) {
	return configs.Etcd{}, nil
}
