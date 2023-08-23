package usecases

import (
	"errors"

	"github.com/bruandreo/order-service/internal/domain"
	"github.com/bruandreo/order-service/internal/repositories"
	"github.com/google/uuid"
)

type NewProduct struct {
	Repository repositories.ProductRepository
}

type NewProductInput struct {
	PizzeriaID  uuid.UUID `json:"pizzeriaId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
}

func (np *NewProduct) NewProductPizzeria(product NewProductInput) (uuid.UUID, error) {
	pizzeria := domain.Pizzeria{ID: product.PizzeriaID}
	newProduct := domain.NewProduct(pizzeria.ID, product.Name, product.Description, product.Price)

	if success := np.Repository.Create(newProduct); success {
		return newProduct.ID, nil
	}
	return uuid.UUID{}, errors.New("fail to create new product")
}
