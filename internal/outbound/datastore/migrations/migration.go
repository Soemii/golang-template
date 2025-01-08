package migrations

import (
	"database/sql"
	"embed"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"io/fs"
)

//go:embed *.sql
var migrations embed.FS

func NewMigration(migrations fs.FS, db *sql.DB) (*migrate.Migrate, error) {
	var m *migrate.Migrate
	var sourceDriver source.Driver
	var databaseDriver database.Driver
	var err error

	databaseDriver, err = postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	sourceDriver, err = iofs.New(migrations, ".")
	if err != nil {
		return nil, err
	}
	m, err = migrate.NewWithInstance("iofs", sourceDriver, "postgres", databaseDriver)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func NewMigrationFiles() fs.FS {
	return migrations
}
