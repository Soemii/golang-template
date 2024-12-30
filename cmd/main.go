package main

import (
	"github.com/Soemii/golang-template/internal/inbound/web"
	"go.uber.org/fx"
	"net/http"
)

func main() {
	fx.New(
		fx.Provide(
			web.NewServer,
			web.NewMux,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
