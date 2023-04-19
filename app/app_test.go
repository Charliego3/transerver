package app

import "testing"

func TestNewApp(t *testing.T) {
	NewApp().Run()
}

func TestDefaultFunc(t *testing.T) {
	var m map[string]struct{}
	m2 := Default(m, make(map[string]struct{}))
	t.Log(m2)
}
