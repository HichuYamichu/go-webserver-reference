package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/HichuYamichu/go-webserver-reference/app/handlers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// App : Application struct
type App struct {
	Router *mux.Router
	DB     *mongo.Database
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
	a.DB = client.Database("Users")

	a.Router = mux.NewRouter()
	a.setRouters()
	a.Adrr = fmt.Sprintf("%s:%s", host, port)
	return a
}

func (a *App) setRouters() {
	api := a.Router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/users", auth(a.handle(handlers.GetUsers))).Methods("GET")
	api.HandleFunc("/user", auth(a.handle(handlers.InsertUser))).Methods("POST")
	api.HandleFunc("/user/{id}", auth(a.handle(handlers.UpdateUser))).Methods("PUT")
	api.HandleFunc("/user/{id}", auth(a.handle(handlers.DeleteUser))).Methods("DELETE")
	a.Router.HandleFunc("/auth", a.handle(handlers.Authenticate)).Methods("GET")
	a.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))
}

type handler func(db *mongo.Database, w http.ResponseWriter, r *http.Request)

func (a *App) handle(h handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(a.DB, w, r)
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
