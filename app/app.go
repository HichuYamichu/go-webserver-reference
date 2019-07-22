package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/HichuYamichu/go-webserver-reference/app/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Router *mux.Router
	DB     *mongo.Database
	Port   string
}

func NewServer(port string, mongoURI string) *App {
	a := &App{}
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}
	CTX, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(CTX)
	if err != nil {
		panic(err)
	}
	a.DB = client.Database("Users")

	a.Router = mux.NewRouter()
	a.setRouters()
	a.Port = port
	return a
}

func (a *App) setRouters() {
	api := a.Router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/users", a.handle(handlers.GetUsers)).Methods("GET")
	api.HandleFunc("/user", a.handle(handlers.InsertUser)).Methods("POST")
	api.HandleFunc("/user/{id}", a.handle(handlers.UpdateUser)).Methods("PUT")
	api.HandleFunc("/user/{id}", a.handle(handlers.DeleteUser)).Methods("DELETE")
	a.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))
}

type handler func(db *mongo.Database, w http.ResponseWriter, r *http.Request)

func (a *App) handle(h handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(a.DB, w, r)
	}
}

func (a *App) Run() {
	srv := &http.Server{
		Handler:      a.Router,
		Addr:         "0.0.0.0:" + a.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Listening on port: %s\n", a.Port)
	log.Fatal(srv.ListenAndServe())
}
