package pkg

type Application struct {
}

func NewApp() *Application {
	return &Application{}
}

func (app *Application) Run() error {
	return nil
}