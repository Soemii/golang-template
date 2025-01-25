package web

import "net/http"

func NewMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		var last http.Handler = route
		for _, middleware := range route.Middlewares() {
			last = middleware.Middleware(last)
		}
		mux.Handle(route.Pattern(), last)
	}
	return mux
}
