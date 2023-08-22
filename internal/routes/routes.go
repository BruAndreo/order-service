package routes

import (
	"github.com/bruandreo/order-service/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Initialize(app *fiber.App, dependencies DependenciesInjector) {
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${method} ${path} ${status} ${latency}",
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]string{
			"status": "healthy",
		})
	})

	productHandler := handlers.ProductsHandler{
		Repository: dependencies.ProductRepository,
	}
	app.Post("/products", productHandler.CreateProduct)
}
