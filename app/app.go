package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	reqCtx "github.com/HichuYamichu/go-webserver-reference/context"
	"github.com/HichuYamichu/go-webserver-reference/handlers/users"
	util "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// App : Application struct
type App struct {
	Router http.Handler
	Adrr   string
}

// NewServer : Initialize new server instance
func NewServer(host, port string) *App {
	a := &App{}
	a.Router = a.setupRouter()
	a.Adrr = fmt.Sprintf("%s:%s", host, port)
	return a
}

func (a *App) setupRouter() http.Handler {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/users", a.handle(users.GetUsers)).Methods("GET")
	// api.HandleFunc("/user", auth(a.handle(handlers.InsertUser))).Methods("POST")
	// api.HandleFunc("/user/{id}", auth(a.handle(handlers.UpdateUser))).Methods("PUT")
	// api.HandleFunc("/user/{id}", auth(a.handle(handlers.DeleteUser))).Methods("DELETE")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))

	allowedHeaders := util.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedOrigins := util.AllowedOrigins([]string{"*"})
	allowedMethods := util.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	h := util.LoggingHandler(os.Stdout, util.CORS(allowedOrigins, allowedHeaders, allowedMethods)(r))
	return h
}

type handler func(ctx *reqCtx.Context)

func (a *App) handle(h handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("content-type", "application/json")
		ctx := reqCtx.CreateContext(w, r)
		h(ctx)
	}
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
