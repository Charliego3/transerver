package configs

import (
	"testing"
)

func TestFetcherTypes(t *testing.T) {
	config, ok := Fetch[Etcd]()
	t.Log(config, ok)
	config, ok = Fetch[Etcd]()
	t.Log(config, ok)
}
