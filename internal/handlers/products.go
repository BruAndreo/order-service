package handlers

import (
	"github.com/bruandreo/order-service/internal/handlers/dto"
	"github.com/bruandreo/order-service/internal/repositories"
	"github.com/bruandreo/order-service/internal/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductsHandler struct {
	Repository repositories.ProductRepository
}

func (ph *ProductsHandler) CreateProduct(c *fiber.Ctx) error {
	pizzeriaId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.MakeErrorDTO(err.Error()))
	}

	bodyDto := dto.NewProductDTO{}

	if err := c.BodyParser(&bodyDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.MakeErrorDTO(err.Error()))
	}

	input := usecases.NewProductInput{
		PizzeriaID:  pizzeriaId,
		Name:        bodyDto.Name,
		Description: bodyDto.Description,
		Price:       bodyDto.Price,
	}

	useCase := usecases.NewProduct{Repository: ph.Repository}
	productId, err := useCase.NewProductPizzeria(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.MakeErrorDTO(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": productId,
	})
}
