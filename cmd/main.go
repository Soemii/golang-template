package main

import (
	"errors"
	"github.com/Soemii/golang-template/internal/inbound/web"
	"github.com/Soemii/golang-template/internal/outbound/datastore"
	"github.com/Soemii/golang-template/internal/outbound/datastore/migrations"
	"github.com/golang-migrate/migrate/v4"
	"go.uber.org/fx"
	"io/fs"
	"net/http"
)

func main() {
	fx.New(
		fx.Provide(
			web.NewServer,
			fx.Annotate(web.NewMux, fx.ParamTags(`group:"route"`)),
			AsRoute(web.NewPrometheusRoute),
		),
		fx.Provide(
			fx.Annotate(migrations.NewMigrationFiles, fx.ResultTags(`name:"migrationFiles"`), fx.As(new(fs.FS))),
			fx.Annotate(migrations.NewMigration, fx.ParamTags(`name:"migrationFiles"`)),
			datastore.NewDB,
		),
		fx.Invoke(func(*http.Server) {}),
		fx.Invoke(func(migration *migrate.Migrate) error {
			err := migration.Up()
			if !errors.Is(err, migrate.ErrNoChange) {
				return err
			}
			return nil
		}),
	).Run()
}

func AsRoute(route any) any {
	return fx.Annotate(route, fx.As(new(web.Route)), fx.ResultTags(`group:"route"`))
}
