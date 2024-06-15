package utils

import (
	"fmt"
	"testing"

	"entgo.io/ent/entc/integration/customid/ent"
)

func TestSingleton(t *testing.T) {
	resp := NewResponse(nil)
	Singleton(resp)

	Call(func(t *ent.Client) error {
		fmt.Printf("%T", t)
		return nil
	})
}
