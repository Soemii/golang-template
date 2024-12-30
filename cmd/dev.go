//go:build dev
// +build dev

package main

import (
	"context"
	"log"
	"os"
)
import "github.com/testcontainers/testcontainers-go/modules/postgres"

func init() {
	postgresContainer, err := postgres.Run(context.Background(),
		"postgres:16-alpine",
		postgres.WithDatabase("test"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		panic(err)
	}
	connectionString, err := postgresContainer.ConnectionString(context.Background())
	if err != nil {
		panic(err)
	}
	err = os.Setenv("DB_DSN", connectionString)
	if err != nil {
		panic(err)
	}
	log.Println("Postgres container started")
}
