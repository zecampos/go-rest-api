package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/zecampos/go-rest-api/internal/transport/http"
)

// App - the struct which contains things like pointes
// to database connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting Up Our App")
	handler := transportHTTP.NewHandler()
	handler.SetupROutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
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
