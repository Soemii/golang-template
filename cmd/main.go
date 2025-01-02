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
			fx.Annotate(web.NewMux, fx.ParamTags(`group:"route"`)),
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

func AsRoute(route any) any {
	return fx.Annotate(route, fx.As(new(web.Route)), fx.ResultTags(`group:"route"`))
}
