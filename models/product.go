package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func ProductsList() ([]Product, error) {
	var products []Product
	result := DB.Find(&products)
	return products, result.Error
}

func ProductByID(id int) (Product, error) {
	var product Product
	result := DB.Take(&product, id)
	return product, result.Error
}

func ProductCreate(product *Product) error {
	return DB.Create(product).Error
}
