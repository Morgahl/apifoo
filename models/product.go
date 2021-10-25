package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func ProductsAll(db *gorm.DB) ([]Product, error) {
	var products []Product
	result := db.Find(&products)
	return products, result.Error
}

func ProductByID(db *gorm.DB, id int) (Product, error) {
	var product Product
	result := db.Take(&product, id)
	return product, result.Error
}

func ProductNew(db *gorm.DB, product *Product) error {
	return db.Create(product).Error
}
