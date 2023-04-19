package configs

import (
	"testing"
)

func TestFetcherTypes(t *testing.T) {
	config, ok := Fetch[EtcdConfig]()
	t.Log(config, ok)
}
