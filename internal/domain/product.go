package domain

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Price       float64
	Active      bool
}

func NewProduct(name, description string, price float64) Product {
	return Product{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Price:       price,
		Active:      true,
	}
}
