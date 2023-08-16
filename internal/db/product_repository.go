package db

import (
	"log"

	"github.com/bruandreo/order-service/internal/domain"
	"github.com/google/uuid"
)

type ProductRepositoryDatabase struct{}

func (prd *ProductRepositoryDatabase) Create(pizzeriaID uuid.UUID, product domain.Product) bool {
	log.Print(pizzeriaID, product)
	// insert db
	return true
}
