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
	http.Server
	DB *gorm.DB
}

// New : Initialize new server instance
func New(host, port string) *App {
	a := &App{}
	a.DB = connectDB()
	a.Addr = fmt.Sprintf("%s:%s", host, port)
	a.Handler = a.setupRouter()
	a.WriteTimeout = 15 * time.Second
	a.ReadTimeout = 15 * time.Second
	return a
}

// Run : Starts the http server
func (a *App) Run() {
	fmt.Printf("Listening on: %s\n", a.Addr)
	log.Fatal(a.ListenAndServe())
}

// Shutdown : Shuts down the http server
func (a *App) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	a.DB.Close()
	if err := a.Server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
}
