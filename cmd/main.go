package main

import (
	"errors"
	"github.com/Soemii/golang-template/internal/inbound/web"
	"github.com/Soemii/golang-template/internal/outbound/datastore"
	"github.com/Soemii/golang-template/internal/outbound/datastore/migrations"
	"github.com/golang-migrate/migrate/v4"
	"github.com/ipfans/fxlogger"
	"go.uber.org/fx"
	"io/fs"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	fx.New(
		fx.WithLogger(fxlogger.WithZerolog(zerolog.New(os.Stdout))),
		fx.Provide(
			web.NewServer,
			web.NewMonitoringHandler,
		),
		fx.Provide(
			fx.Annotate(migrations.NewMigrationFiles, fx.ResultTags(`name:"migrationFiles"`), fx.As(new(fs.FS))),
			fx.Annotate(migrations.NewMigration, fx.ParamTags(`name:"migrationFiles"`)),
			datastore.NewDB,
		),
		fx.Invoke(web.SetupRoutes),
		fx.Invoke(func(migration *migrate.Migrate) error {
			err := migration.Up()
			if !errors.Is(err, migrate.ErrNoChange) {
				return err
			}
			return nil
		}),
	).Run()
}
