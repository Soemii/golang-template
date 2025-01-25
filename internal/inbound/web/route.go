package web

import "net/http"

type Route interface {
	Pattern() string
	Middlewares() []Middleware
	http.Handler
}
