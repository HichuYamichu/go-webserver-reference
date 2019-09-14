package app

import (
	"log"
	"net/http"

	"github.com/hichuyamichu/go-webserver-reference/controllers/base"
	"github.com/hichuyamichu/go-webserver-reference/controllers/user"
)

func (a *App) setupRouter() http.Handler {
	r := http.NewServeMux()
	userCont := user.NewUserController(a.DB)
	userCont.Use(first)
	userCont.Use(second)
	r.HandleFunc("/api/user/create", userCont.Run(userCont.InsertUser))
	r.HandleFunc("/api/user/read", userCont.Run(userCont.GetUsers, third))
	r.HandleFunc("/api/user/update", userCont.Run(userCont.UpdateUser))
	r.HandleFunc("/api/user/delete", userCont.Run(userCont.DeleteUser))
	r.Handle("/", http.FileServer(http.Dir("web")))

	return r
}

// middleware example
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
