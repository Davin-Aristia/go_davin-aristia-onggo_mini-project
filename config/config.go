package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	JWT_KEY     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load env file. Err: %s", err)
	}

	JWT_KEY = os.Getenv("JWT_KEY")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_PORT = os.Getenv("DB_PORT")
	DB_HOST = os.Getenv("DB_HOST")
	DB_NAME = os.Getenv("DB_NAME")
}
