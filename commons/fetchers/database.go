package fetchers

import "github.com/transerver/app/configs"

type Database struct{}

func (f *Database) Fetch() (configs.Database, error) {
	return configs.Database{}, nil
}
