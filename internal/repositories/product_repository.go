package repositories

import (
	"github.com/bruandreo/order-service/internal/domain"
)

type ProductRepository interface {
	Create(product domain.Product) bool
}
