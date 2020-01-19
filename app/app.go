package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	api := e.Group("/api")
	userService := api.Group("/user")
	userService.GET("/:id", handlers.FindUser)
	userService.GET("/", handlers.FindUsers)
	userService.POST("", handlers.CreateUser)
	userService.PUT("/", handlers.UpdateUser)
	userService.DELETE("/:id", handlers.DeleteUser)

	return e
}
