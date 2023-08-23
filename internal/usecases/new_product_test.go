package usecases_test

import (
	"testing"

	"github.com/bruandreo/order-service/internal/domain"
	"github.com/bruandreo/order-service/internal/usecases"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type ProductRepositoryMock struct {
	CreateReturn bool
}

func (rm *ProductRepositoryMock) Create(product domain.Product) bool {
	return rm.CreateReturn
}

func TestNewProduct(t *testing.T) {

	t.Run("Should create new product and return ID", func(t *testing.T) {
		product := usecases.NewProductInput{
			PizzeriaID:  uuid.New(),
			Name:        "test",
			Description: "test",
			Price:       45.67,
		}

		repoMock := ProductRepositoryMock{CreateReturn: true}
		useCase := usecases.NewProduct{Repository: &repoMock}

		newProductID, err := useCase.NewProductPizzeria(product)

		assert.Nil(t, err)
		assert.NotEmpty(t, newProductID)
	})

	t.Run("Shouldn't create new product and return Error", func(t *testing.T) {
		product := usecases.NewProductInput{
			PizzeriaID:  uuid.New(),
			Name:        "test",
			Description: "test",
			Price:       45.67,
		}

		repoMock := ProductRepositoryMock{CreateReturn: false}
		useCase := usecases.NewProduct{Repository: &repoMock}

		newProductID, err := useCase.NewProductPizzeria(product)

		assert.NotNil(t, err)
		assert.Error(t, err, "fail to insert new product")
		assert.Empty(t, newProductID)
	})
}
