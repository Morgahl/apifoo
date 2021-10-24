package models

import "gorm.io/gorm"

func INIT(db *gorm.DB) error {
	return db.AutoMigrate(
		&Product{},
	)
}
