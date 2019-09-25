package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/hichuyamichu/go-webserver-reference/controllers/user"
	"github.com/hichuyamichu/go-webserver-reference/models"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
)

// App : Application struct
type App struct {
	srv *http.Server
	reg *prometheus.Registry
	db  *gorm.DB
}

// New : Initialize new server instance
func New(host, port string) *App {
	a := &App{}
	a.db = a.connectDB()
	a.reg = prometheus.NewRegistry()
	a.reg.MustRegister(prometheus.NewGoCollector())
	a.reg.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	a.srv = &http.Server{}
	a.srv.Addr = fmt.Sprintf("%s:%s", host, port)
	a.srv.Handler = a.setupHandler()
	a.srv.WriteTimeout = 15 * time.Second
	a.srv.ReadTimeout = 15 * time.Second
	return a
}

func (a *App) setupHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)

	userCont := user.NewController(a.db)

	r.Route("/api/user", func(r chi.Router) {
		r.Get("/", handle(userCont.GetAllUsers))
		r.Post("/", handle(userCont.InsertUser))

		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", handle(userCont.GetUser))
			r.Put("/", handle(userCont.UpdateUser))
			r.Delete("/", handle(userCont.DeleteUser))
		})
	})

	r.Handle("/metrics", promhttp.HandlerFor(a.reg, promhttp.HandlerOpts{}))
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

func (a *App) connectDB() *gorm.DB {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")

	uri := fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable", host, user, dbname, password)
	db, err := gorm.Open("postgres", uri)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})

	return db
}

// Run : Starts the app
func (a *App) Run() {
	log.Printf("Listening on: http://%s\n", a.srv.Addr)
	log.Fatal(a.srv.ListenAndServe())
}

// Shutdown : Stops the app
func (a *App) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := a.srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
}
