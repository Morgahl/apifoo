package main

import (
	"os"
	"strings"
)

func GetConfig() Config {
	return Config{
		LoggerConfig: LoggerConfig{
			Debug:       getEnvOrDefaultBool("LOGGER_DEBUG", false),
			PrettyPrint: getEnvOrDefaultBool("LOGGER_PRETTY", false),
		},
		DatabaseName: getEnvOrDefaultString("DATABASE_NAME", "development"),
		ApplicationConfig: ApplicationConfig{
			Port: getEnvOrDefaultString("PORT", ":8080"),
		},
	}
}

type Config struct {
	LoggerConfig      LoggerConfig
	DatabaseName      string
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
