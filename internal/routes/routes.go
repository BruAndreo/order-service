package routes

import (
	"github.com/bruandreo/order-service/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func Initialize(app *fiber.App, dependencies DependenciesInjector) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]string{
			"status": "healthy",
		})
	})

	productHandler := handlers.ProductsHandler{
		Repository: dependencies.ProductRepository,
	}
	app.Post("/pizzeria/:id/products", productHandler.CreateProduct)
}
