package main

import (
	"fmt"
	"log"
	transportHTTP "github.com/devgit072/production-ready-rest-api-in-go/internal/pkg/transport/http"
	"net/http"
)

const port = 7000

type App struct {

}

func (app *App) Run() error {
	log.Println("Starting the server")
	h := transportHTTP.NewHandler()
	h.CreateRoutes()
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), h.Router); err != nil {
		log.Fatalf("Error: %s", err.Error())
		return err
	}
	log.Println("Server is succesfully started on port: ", port)
	return nil
}

func main() {
	fmt.Println("A production ready rest api in go")
	app := App{}
	if err := app.Run(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
