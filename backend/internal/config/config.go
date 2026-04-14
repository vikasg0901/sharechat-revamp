package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds all runtime configuration loaded from environment variables.
type Config struct {
	// HTTP server
	Port string

	// Database (optional — omit DB_URL to skip DB health check)
	DBURL string

	// Redis (optional — omit REDIS_URL to skip Redis health check)
	RedisURL string

	// Observability
	LogLevel string // debug, info, warn, error
	Env      string // development, staging, production

	// Rate limiting
	RateLimitRPM int
}

// Load reads configuration from environment variables and returns a Config.
func Load() (*Config, error) {
	cfg := &Config{
		Port:     os.Getenv("PORT"),
		DBURL:    os.Getenv("DB_URL"),
		RedisURL: os.Getenv("REDIS_URL"),
		LogLevel: os.Getenv("LOG_LEVEL"),
		Env:      os.Getenv("ENV"),
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	if cfg.LogLevel == "" {
		cfg.LogLevel = "info"
	}
	if cfg.Env == "" {
		cfg.Env = "development"
	}

	if rpmStr := os.Getenv("RATE_LIMIT_RPM"); rpmStr != "" {
		rpm, err := strconv.Atoi(rpmStr)
		if err != nil {
			return nil, fmt.Errorf("RATE_LIMIT_RPM must be a valid integer: %w", err)
		}
		cfg.RateLimitRPM = rpm
	} else {
		cfg.RateLimitRPM = 100
	}

	return cfg, nil
}
