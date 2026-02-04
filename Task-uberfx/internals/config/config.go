package config

import (
	"fmt"
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
	APP_PORT  string
	DB_DRIVER string

	//mysql

	MY_DRIVER  string
	MY_HOST    string
	MY_USER    string
	MY_PASS    string
	MY_DB_NAME string
	MY_PORT    string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found, relying on system environment")
	}

	cfg := &Config{
		DB_DRIVER: os.Getenv("DB_DRIVER"),
		DB_HOST:   os.Getenv("DB_HOST"),
		DB_USER:   os.Getenv("DB_USER"),
		DB_PASS:   os.Getenv("DB_PASS"),
		DB_NAME:   os.Getenv("DB_NAME"),
		DB_PORT:   os.Getenv("DB_PORT"),
		APP_PORT:  os.Getenv("APP_PORT"),

		// MySQL
		MY_DRIVER:  os.Getenv("MY_DRIVER"),
		MY_HOST:    os.Getenv("MY_HOST"),
		MY_USER:    os.Getenv("MY_USER"),
		MY_PASS:    os.Getenv("MY_PASS"),
		MY_DB_NAME: os.Getenv("MY_DB_NAME"),
		MY_PORT:    os.Getenv("MY_PORT"), // must match .env exactly
	}

	fmt.Printf("Loaded config: %+v\n", cfg) // <- debug
	return cfg, nil
}
