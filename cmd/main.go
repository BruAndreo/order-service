package main

import (
	"github.com/bruandreo/order-service/internal/config"
	"github.com/bruandreo/order-service/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.Initialize()

	app := fiber.New(fiber.Config{
		AppName: "Order Service API",
	})

	routes.Initialize(app)

	app.Listen(":" + config.AppPort)
}
