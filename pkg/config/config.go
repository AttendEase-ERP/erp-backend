package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL    string
	Port           string
	ClerkSecretKey string
	Env            string
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")

	cfg := &Config{
		DatabaseURL:    getEnv("DATABASE_URL", ""),
		Port:           getEnv("PORT", "5000"),
		ClerkSecretKey: getEnv("CLERK_SECRET_KEY", ""),
		Env:            getEnv("ENV", "development"),
	}

	if cfg.DatabaseURL == "" {
		log.Fatalf("DATABASE_URL is required but not set")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
