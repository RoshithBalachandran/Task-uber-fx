package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST   string
	DB_USER   string
	DB_PASS   string
	DB_NAME   string
	DB_PORT   string
	MYSQL_DSN string
	APP_PORT  string
}

func Load() *Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("failed to load env")
	}
	return &Config{
		DB_HOST:   os.Getenv("DB_HOST"),
		DB_USER:   os.Getenv("DB_USER"),
		DB_PASS:   os.Getenv("DB_PASS"),
		DB_NAME:   os.Getenv("DB_NAME"),
		DB_PORT:   os.Getenv("DB_PORT"),
		MYSQL_DSN: os.Getenv("MYSQL_DSN"),
		APP_PORT:  os.Getenv("APP_PORT"),
	}
}
