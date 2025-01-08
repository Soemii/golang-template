package datastore

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
)

func NewDB() (*sql.DB, error) {
	d, err := sql.Open("pgx", os.Getenv("DB_DSN"))
	if err != nil {
		return nil, err
	}
	if err := d.Ping(); err != nil {
		return nil, err
	}
	return d, nil
}
