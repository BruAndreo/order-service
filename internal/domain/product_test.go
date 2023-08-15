package domain_test

import (
	"testing"

	"github.com/bruandreo/order-service/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewProductWithNameDescriptionPrice(t *testing.T) {
	expectedName := "test"
	expectedDescription := "test"
	expectedPrice := 45.67

	product := domain.NewProduct(expectedName, expectedDescription, expectedPrice)

	assert.NotEmpty(t, product.ID)
	assert.Equal(t, expectedName, product.Name)
	assert.Equal(t, expectedDescription, product.Description)
	assert.Equal(t, expectedPrice, product.Price)
	assert.True(t, product.Active)
}
