package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenRouterAPIKey string

	PostgresHost   string
	PostgresPort   string
	PostgresUser   string
	PostgresPass   string
	PostgresDBName string
}

func Init() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		OpenRouterAPIKey: os.Getenv("OPEN_ROUTER_API_KEY"),
		PostgresHost:     os.Getenv("PostgresHost"),
		PostgresPort:     os.Getenv("PostgresPort"),
		PostgresUser:     os.Getenv("PostgresUser"),
		PostgresPass:     os.Getenv("PostgresPass"),
		PostgresDBName:   os.Getenv("PostgresDBName"),
	}, nil
}
