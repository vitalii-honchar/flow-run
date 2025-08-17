package config

import (
	"flow-run/internal/flowrun/infra/database"
	"flow-run/internal/lib/logger"
	"flow-run/internal/lib/validator"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	*database.DatabaseConfig
	ServerPort string `validate:"required,numeric,min=1,max=65535"`
	ServerHost string `validate:"required,ip"`
}

func FromEnv() (*Config, error) {
	_ = godotenv.Load()

	config := &Config{
		DatabaseConfig: &database.DatabaseConfig{
			URL:             os.Getenv("DATABASE_URL"),
			MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 10),
			ConnMaxLifetime: getEnvAsDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute),
		},
		ServerPort: getEnvWithDefault("SERVER_PORT", "8080"),
		ServerHost: getEnvWithDefault("SERVER_HOST", "0.0.0.0"),
	}

	return validator.Struct(config)
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		logger.Log.Warningf("Invalid integer value for %s: %s, using default: %d", key, value, defaultValue)
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
		logger.Log.Warningf("Invalid duration value for %s: %s, using default: %s", key, value, defaultValue)
	}
	return defaultValue
}
