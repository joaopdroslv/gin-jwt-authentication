package config

import (
	"os"
	"strconv"
)

type Env struct {
	HTTPPort          string
	DockerDatabaseURL string
	JWTSecret         string
	JWTTTL            int64
}

func NewEnv() *Env {

	return &Env{
		HTTPPort:          getEnv("HTTP_PORT", "8080"),
		DockerDatabaseURL: getEnv("DOCKER_DATABASE_URL", "mysql"),
		JWTSecret:         getEnv("JWT_SECRET", ""),
		JWTTTL:            getEnvInt64("JWT_TTL", 3600),
	}
}

func getEnv(key, fallback string) string {

	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func getEnvInt64(key string, fallback int64) int64 {

	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return fallback
	}

	return parsed
}
