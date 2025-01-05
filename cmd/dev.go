//go:build dev
// +build dev

package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(os.Getenv("DB_DSN"))
}
