package web

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func NewPrometheusRoute() Route {
	return prometheusRoute{}
}

type prometheusRoute struct{}

func (r prometheusRoute) Middlewares() []Middleware {
	return []Middleware{}
}

func (prometheusRoute) Pattern() string {
	return "/metrics"
}

func (prometheusRoute) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	promhttp.Handler().ServeHTTP(writer, request)
}
