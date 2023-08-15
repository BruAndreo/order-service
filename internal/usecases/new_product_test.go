package usecases_test

import (
	"testing"

	"github.com/bruandreo/order-service/internal/usecases"
	"github.com/stretchr/testify/assert"
)

func TestInsertNewProduct(t *testing.T) {
	// pizzeria
	// product

	newProductID := usecases.NewProductPizzeria()

	assert.NotEmpty(t, newProductID)
}
