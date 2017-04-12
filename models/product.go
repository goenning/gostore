package models

import "math/big"

// Product is what we sell on our store
type Product struct {
	Title       string
	Description string
	Price       *big.Float
}

// NewProduct returns a new Product instance
func NewProduct(title, description string, price *big.Float) *Product {
	return &Product{
		Title:       title,
		Description: description,
		Price:       price,
	}
}
