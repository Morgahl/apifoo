package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Logger(config LoggerConfig) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if config.PrettyPrint {
		log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
}

type LoggerConfig struct {
	Debug       bool
	PrettyPrint bool
}
