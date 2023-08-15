package usecases

import (
	"github.com/bruandreo/order-service/internal/domain"
	"github.com/google/uuid"
)

type ProductBasic struct {
	Name        string
	Description string
	Price       float64
}

func NewProductPizzeria(pizzeria domain.Pizzeria, product ProductBasic) uuid.UUID {
	newProduct := domain.NewProduct(product.Name, product.Description, product.Price)

	// chama repo para cadastrar

	return newProduct.ID
}
