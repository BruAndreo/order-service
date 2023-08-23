package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	PizzeriaID  uuid.UUID
	Name        string
	Description string
	Price       float64
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProduct(pizzeriaId uuid.UUID, name, description string, price float64) Product {
	return Product{
		ID:          uuid.New(),
		PizzeriaID:  pizzeriaId,
		Name:        name,
		Description: description,
		Price:       price,
		Active:      true,
	}
}
