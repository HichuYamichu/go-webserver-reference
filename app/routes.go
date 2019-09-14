package app

import (
	"log"
	"net/http"

	"github.com/hichuyamichu/go-webserver-reference/controllers/base"
	"github.com/hichuyamichu/go-webserver-reference/controllers/user"
)

func (a *App) setupRouter() http.Handler {
	r := http.NewServeMux()
	userCont := user.NewUserController()
	userCont.Use(first)
	userCont.Use(second)
	r.HandleFunc("/api/users", userCont.Run(userCont.GetUsers, third))
	// api.HandleFunc("/user", auth(a.handle(handlers.InsertUser))).Methods("POST")
	// api.HandleFunc("/user/{id}", auth(a.handle(handlers.UpdateUser))).Methods("PUT")
	// api.HandleFunc("/user/{id}", auth(a.handle(handlers.DeleteUser))).Methods("DELETE")
	// r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))

	return r
}

func first(next base.Handler) base.Handler {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		log.Printf("First started")
		next(w, r)
		log.Printf("First finished")
		return 200, nil
	}
}

func second(next base.Handler) base.Handler {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		log.Printf("Second started")
		next(w, r)
		log.Printf("Second finished")
		return 200, nil
	}
}

func third(next base.Handler) base.Handler {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		log.Printf("Third started")
		next(w, r)
		log.Printf("Third finished")
		return 200, nil
	}
}
