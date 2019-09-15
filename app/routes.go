package app

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/hichuyamichu/go-webserver-reference/controllers/user"
)

func (a *App) setupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	userCont := user.NewController(a.DB)

	r.Route("/api/user", func(r chi.Router) {
		r.Get("/", handle(userCont.GetAllUsers))
		r.Post("/", handle(userCont.InsertUser))

		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", handle(userCont.GetUser))
			r.Put("/", handle(userCont.UpdateUser))
			r.Delete("/", handle(userCont.DeleteUser))
		})
	})

	r.Handle("/", http.FileServer(http.Dir("web")))

	return r
}

type handler func(w http.ResponseWriter, r *http.Request) (int, error)

func handle(h handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if code, err := h(w, r); err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(code), code)
		}
	}
}
