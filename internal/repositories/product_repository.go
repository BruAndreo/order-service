package repositories

import (
	"github.com/bruandreo/order-service/internal/domain"
	"github.com/google/uuid"
)

type ProductRepository interface {
	Create(pizzeriaID uuid.UUID, product domain.Product) bool
}
