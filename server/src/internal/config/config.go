package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host      string
	DbUrl     string
	JwtSecret string
}

func Load() (*Config, error) {
	godotenv.Load()
	host, ok := os.LookupEnv("HOST")
	if !ok {
		return nil, errors.New("'HOST' not provided in env")
	}
	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return nil, errors.New("'DB_NAME' not provided in env")
	}
	dbUsername, ok := os.LookupEnv("DB_USERNAME")
	if !ok {
		return nil, errors.New("'DB_USERNAME' not provided in env")
	}
	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return nil, errors.New("'DB_PASSWORD' not provided in env")
	}
	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return nil, errors.New("'DB_HOST' not provided in env")
	}
	dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUsername, dbPassword, dbHost, dbName)
	jwtSecret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return nil, errors.New("'JWT_SECRET' not provided in env")
	}
	return &Config{
		Host:      host,
		DbUrl:     dbURL,
		JwtSecret: jwtSecret,
	}, nil
}
