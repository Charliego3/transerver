package main

import "github.com/transerver/app"

func main() {
	app := app.NewApp(app.WithAddr("tcp", ":8080"))
	app.Run()
}
