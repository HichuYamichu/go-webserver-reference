package app

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// App : Application struct
type App struct {
	Router http.Handler
	Adrr   string
}

// New : Initialize new server instance
func New(host, port string) *App {
	a := &App{}
	a.Router = a.setupRouter()
	a.Adrr = fmt.Sprintf("%s:%s", host, port)
	return a
}

// Run : Starts the http server
func (a *App) Run() {
	srv := &http.Server{
		Handler:      a.Router,
		Addr:         a.Adrr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Listening on: %s\n", a.Adrr)
	log.Fatal(srv.ListenAndServe())
}
