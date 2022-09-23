package main

func main() {
	app, cleanup, err := wireApp()
	if err != nil {
		panic(err)
	}

	defer cleanup()
	app.Run()
}
