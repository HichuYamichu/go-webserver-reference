package context

import (
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func CreateContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{W: w, R: r}
}