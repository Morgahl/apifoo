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
	if err := models.INIT(config.DatabaseName); err != nil {
		log.Fatal().Err(err).Msg("error setting up database")
	}

	// Setup API
	log.Info().Err(Router().Run(config.ApplicationConfig.Port)).Msg("Server shutdown")
}

type ApplicationConfig struct {
	Port string
}
