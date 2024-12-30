package web

import (
	"go.uber.org/fx"
	"net/http"
)

func NewServer(lifecycle fx.Lifecycle, mux *http.ServeMux) *http.Server {
	serv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	lifecycle.Append(fx.StartStopHook(func() {
		go serv.ListenAndServe()
	}, func() error {
		return serv.Close()
	}))
	return serv
}
