package domain_test

import (
	"testing"

	"github.com/bruandreo/order-service/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	t.Run("Should Create new Product when Name, Description and Price are passed", func(t *testing.T) {
		expectedPizzeriaId := uuid.MustParse("9ea3ef09-f05c-498d-9e25-152d4df438ac")
		expectedName := "test"
		expectedDescription := "test"
		expectedPrice := 45.67

		product := domain.NewProduct(expectedPizzeriaId, expectedName, expectedDescription, expectedPrice)

		assert.NotEmpty(t, product.ID)
		assert.Equal(t, expectedName, product.Name)
		assert.Equal(t, expectedDescription, product.Description)
		assert.Equal(t, expectedPrice, product.Price)
		assert.True(t, product.Active)
	})
}
