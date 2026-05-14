package config

import "os"

type Config struct {
	AppPort        string
	DatabaseURL    string
	RedisAddr      string
	RedisURL       string
	JWTSecret      string
	AllowedOrigins string
}

func Load() Config {
	return Config{
		AppPort: getEnv("APP_PORT", getEnv("PORT", "8080")),
		DatabaseURL: getEnv(
			"DATABASE_URL",
			"postgres://figureshelf:figureshelf123@localhost:5433/figureshelf_db?sslmode=disable",
		),
		RedisAddr:      getEnv("REDIS_ADDR", "localhost:6380"),
		RedisURL:       getEnv("REDIS_URL", ""),
		JWTSecret:      getEnv("JWT_SECRET", "figureshelf-dev-secret"),
		AllowedOrigins: getEnv("ALLOWED_ORIGINS", "http://localhost:5173"),
	}
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}