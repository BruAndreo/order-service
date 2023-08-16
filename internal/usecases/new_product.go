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

type ProductBasic struct {
	Name        string
	Description string
	Price       float64
}

func (np *NewProduct) NewProductPizzeria(pizzeria domain.Pizzeria, product ProductBasic) (uuid.UUID, error) {
	newProduct := domain.NewProduct(product.Name, product.Description, product.Price)

	if success := np.Repository.Create(pizzeria.ID, newProduct); success {
		return newProduct.ID, nil
	}
	return uuid.UUID{}, errors.New("fail to create new product")
}
