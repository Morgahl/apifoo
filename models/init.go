package models

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	User   string
	Pass   string
	Host   string
	Port   int
	DBName string
}

func INIT(config Config) {
	var err error
	DB, err = gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		config.Host, config.User, config.Pass, config.DBName, config.Port,
	)), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to setup database")
	}
	if err = migrate(); err != nil {
		log.Fatal().Err(err).Msg("failed to migrate")
	}
}

func migrate() error {
	return DB.AutoMigrate(
		&Product{},
	)
}
