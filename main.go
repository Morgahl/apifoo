package main

import (
	"github.com/curlymon/gormfoo/models"
	"github.com/rs/zerolog/log"
)

func main() {
	config := GetConfig()

	// Setup Logger
	Logger(config.LoggerConfig)

	// Setup Database
	db, err := Database(config.DatabaseConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to setup database")
	}

	if err := models.INIT(db); err != nil {
		log.Fatal().Err(err).Msg("failed to migrate models")
	}

	// Setup API
	log.Info().Err(Router(db).Run(config.ApplicationConfig.Port)).Msg("Server shutdown")
}

type ApplicationConfig struct {
	Port string
}
