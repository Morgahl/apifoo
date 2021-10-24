package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database(config DatabaseConfig) (*gorm.DB, error) {
	return gorm.Open(
		postgres.Open(
			fmt.Sprintf(
				"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
				config.Host,
				config.User,
				config.Pass,
				config.DBName,
				config.Port,
			),
		),
		&gorm.Config{})
}

type DatabaseConfig struct {
	User   string
	Pass   string
	Host   string
	Port   int
	DBName string
}
