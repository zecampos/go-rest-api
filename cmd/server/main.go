package main

import "fmt"

// App - the struct which contains things like pointes
// to database connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting Up Our App")
	return nil
}

func main() {
	fmt.Println("GO Rest API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our Rest API")
		fmt.Println(err)
	}
}
