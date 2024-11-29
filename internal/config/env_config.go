package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
	DBPort     string
}

func LoadDBConfig() DBConfig {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	return DBConfig{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
	}

}
