package web

import "net/http"

type Route interface {
	Pattern() string
	http.Handler
}
