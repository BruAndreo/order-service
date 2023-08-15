package main

import (
	"github.com/bruandreo/order-service/internal/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.Initialize()

	app := fiber.New(fiber.Config{
		AppName: "Order Service API",
	})

	app.Listen(":" + config.AppPort)
}
