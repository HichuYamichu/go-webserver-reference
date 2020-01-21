package app

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/hichuyamichu/go-webserver-reference/handlers"
)

// New creates new app instance
func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.AddTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	api := e.Group("/api")
	userService := api.Group("/user")
	userService.GET("/:id", handlers.FindUser)
	userService.GET("/", handlers.FindUsers)
	userService.POST("", handlers.CreateUser)
	userService.PUT("/", handlers.UpdateUser)
	userService.DELETE("/:id", handlers.DeleteUser)

	return e
}
