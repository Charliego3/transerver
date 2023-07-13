package app

import "testing"

func TestNewApp(t *testing.T) {
	NewApp().Run()
}

func TestDefaultFunc(t *testing.T) {
	var arr []string
	t.Log(append([]string{"first"}, arr...))
}
