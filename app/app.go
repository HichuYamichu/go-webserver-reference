package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

// App : Application struct
type App struct {
	Server *http.Server
	Router http.Handler
	DB     *gorm.DB
	Adrr   string
}

// New : Initialize new server instance
func New(host, port string) *App {
	a := &App{}
	a.DB = connectDB()
	a.Router = a.setupRouter()
	a.Adrr = fmt.Sprintf("%s:%s", host, port)
	a.Server = &http.Server{
		Handler:      a.Router,
		Addr:         a.Adrr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return a
}

// Run : Starts the http server
func (a *App) Run() {
	fmt.Printf("Listening on: %s\n", a.Adrr)
	log.Fatal(a.Server.ListenAndServe())
}

func (a *App) Shutdown() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	a.DB.Close()
	if err := a.Server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
}
