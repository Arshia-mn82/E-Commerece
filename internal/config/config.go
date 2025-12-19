package config

import "os"

type Config struct {
	AppPort   string
	DBDsn     string
	JWTSecret string
}

func Load() *Config {
	return &Config{
		AppPort:   getEnv("APP_PORT", "8080"),
		DBDsn:     getEnv("DB_DSN", ""),
		JWTSecret: getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
