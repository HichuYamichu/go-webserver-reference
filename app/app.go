package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	appCtx "github.com/HichuYamichu/go-webserver-reference/app/context"
	"github.com/HichuYamichu/go-webserver-reference/app/handlers"
	"github.com/dgrijalva/jwt-go"
	util "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// App : Application struct
type App struct {
	Router http.Handler
	Ctx    *appCtx.Context
	Adrr   string
}

// NewServer : Initialize new server instance
func NewServer(host, port, mongoURI string) *App {
	a := &App{}
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	db := client.Database("Users")

	a.Ctx = appCtx.NewContext(db)
	a.Router = a.setupRouter()
	a.Adrr = fmt.Sprintf("%s:%s", host, port)
	return a
}

func (a *App) setupRouter() http.Handler {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/users", auth(a.handle(handlers.GetUsers))).Methods("GET")
	api.HandleFunc("/user", auth(a.handle(handlers.InsertUser))).Methods("POST")
	api.HandleFunc("/user/{id}", auth(a.handle(handlers.UpdateUser))).Methods("PUT")
	api.HandleFunc("/user/{id}", auth(a.handle(handlers.DeleteUser))).Methods("DELETE")
	r.HandleFunc("/auth", a.handle(handlers.Authenticate)).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))

	allowedHeaders := util.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedOrigins := util.AllowedOrigins([]string{"*"})
	allowedMethods := util.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	h := util.LoggingHandler(os.Stdout, util.CORS(allowedOrigins, allowedHeaders, allowedMethods)(r))
	return h
}

type handler func(ctx *appCtx.Context, w http.ResponseWriter, r *http.Request) *handlers.AppError

func (a *App) handle(h handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("content-type", "application/json")
		if err := h(a.Ctx, w, r); err != nil {
			fmt.Println(err)
			http.Error(w, err.Msg, err.Code)
		}
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

var secret = os.Getenv("SECRET")

func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tokenString := c.Value
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
