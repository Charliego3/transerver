package main

import "github.com/google/wire"

var providerSet = wire.NewSet()

func main() {
	app, cleanup, err := wireApp()
	if err != nil {
		panic(err)
	}

	defer cleanup()
	app.Run()
}
