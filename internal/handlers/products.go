package handlers

import (
	"github.com/bruandreo/order-service/internal/repositories"
	"github.com/bruandreo/order-service/internal/usecases"
	"github.com/gofiber/fiber/v2"
)

type ProductsHandler struct {
	Repository repositories.ProductRepository
}

func (ph *ProductsHandler) CreateProduct(c *fiber.Ctx) error {
	useCase := usecases.NewProduct{
		Repository: ph.Repository,
	}
	input := usecases.NewProductInput{}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	productId, err := useCase.NewProductPizzeria(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": productId,
	})
}
