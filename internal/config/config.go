package config

import (
	"os"
	"time"
)

type Config struct {
	Port           string
	DatabaseURL    string
	JWTSecret      string
	JWTExpiry      time.Duration
}

func Load() *Config {
	return &Config{
		Port:           getEnv("PORT", "8000"),
		DatabaseURL:    getEnv("DATABASE_URL", "postgres://postgres:password@localhost:5432/fastline?sslmode=disable"),
		JWTSecret:      getEnv("JWT_SECRET","2499d600f7df1ad1b42f8e1b34938861"),
		JWTExpiry:      time.Duration(time.Minute * 15),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
