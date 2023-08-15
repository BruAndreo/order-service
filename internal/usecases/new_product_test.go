package usecases_test

import (
	"testing"

	"github.com/bruandreo/order-service/internal/domain"
	"github.com/bruandreo/order-service/internal/usecases"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsertNewProduct(t *testing.T) {
	pizzeria := domain.Pizzeria{ID: uuid.New(), Name: "Pizza Test"}
	product := usecases.ProductBasic{
		Name:        "test",
		Description: "test",
		Price:       45.67,
	}

	newProductID := usecases.NewProductPizzeria(pizzeria, product)

	assert.NotEmpty(t, newProductID)
}
