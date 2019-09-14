package base

import (
	"log"
	"net/http"
)

type Controller struct {
	Middleware []Middleware
}

type Middleware func(Handler) Handler

type Handler func(w http.ResponseWriter, r *http.Request) (int, error)

func (c *Controller) Use(mw Middleware) {
	c.Middleware = append(c.Middleware, mw)
}

func (c *Controller) Run(h Handler, mw ...Middleware) http.HandlerFunc {
	last := h
	c.Middleware = append(c.Middleware, mw...)
	for i := len(c.Middleware) - 1; i >= 0; i-- {
		last = c.Middleware[i](last)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if code, err := last(w, r); err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(code), code)
		}
	}
}
