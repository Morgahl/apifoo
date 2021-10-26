package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Product is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Product struct {
	ID        uuid.UUID `json:"-" db:"id"`
	Code      string    `json:"code" db:"code"`
	Price     float64   `json:"price" db:"price"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func ProductList() (products []Product, err error) {
	err = DB.All(&products)
	return
}

func ProductByID(id string) (product Product, err error) {
	err = DB.Find(&product, id)
	return
}

func ProductCreate(product Product) error {
	return DB.Create(product)
}
