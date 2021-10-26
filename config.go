package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/curlymon/gormfoo/models"
)

func GetConfig() Config {
	return Config{
		LoggerConfig: LoggerConfig{
			Debug:       getEnvOrDefaultBool("LOGGER_DEBUG", false),
			PrettyPrint: getEnvOrDefaultBool("LOGGER_PRETTY", false),
		},
		DatabaseConfig: models.Config{
			User:   getEnvOrDefaultString("POSTGRES_USER", ""),
			Pass:   getEnvOrDefaultString("POSTGRES_PASSWORD", ""),
			Host:   getEnvOrDefaultString("POSTGRES_HOST", ""),
			Port:   getEnvOrDefaultInt("POSTGRES_PORT", 5432),
			DBName: getEnvOrDefaultString("POSTGRES_DB", ""),
		},
		ApplicationConfig: ApplicationConfig{
			Port: getEnvOrDefaultString("PORT", ":8080"),
		},
	}
}

type Config struct {
	LoggerConfig      LoggerConfig
	DatabaseConfig    models.Config
	ApplicationConfig ApplicationConfig
}

func getEnvOrDefaultString(key, def string) string {
	if env := os.Getenv(key); env != "" {
		return env
	}

	return def
}

func getEnvOrDefaultBool(key string, def bool) bool {
	if env := os.Getenv(key); env != "" {
		switch strings.ToLower(env) {
		case "yes", "true", "1":
			return true
		default:
			return false
		}
	}

	return def
}

func getEnvOrDefaultInt(key string, def int) int {
	if env := os.Getenv(key); env != "" {
		if i, err := strconv.Atoi(env); err == nil {
			return i
		}
	}

	return def
}
