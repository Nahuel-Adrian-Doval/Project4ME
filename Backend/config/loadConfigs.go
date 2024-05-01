package config

import (
	"os"

	"github.com/Nahuel-Adrian-Doval/Project4ME/Backend/models"
)

func LoadDBConfig() models.DBConfig {
	return models.DBConfig{
		DATABASE_PORT: os.Getenv("DATABASE_PORT"),

		DATABASE_HOST:     os.Getenv("DATABASE_HOST"),
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),

		DATABASE_SSLMODE:  os.Getenv("DATABASE_SSLMODE"),
		DATABASE_TIMEZONE: os.Getenv("DATABASE_TIMEZONE"),
	}
}

func LoadServerConfig() models.ServerConfig {
	return models.ServerConfig{
		SERVER_PORT:   os.Getenv("SERVER_PORT"),
		SERVER_PREFIX: os.Getenv("SERVER_PREFIX"),
	}
}
