package main

import (
	"github.com/bruandreo/order-service/internal/config"
	"github.com/bruandreo/order-service/internal/db"
	"github.com/bruandreo/order-service/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func ApiCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Execute API",
		Run: func(cmd *cobra.Command, args []string) {
			ExecuteApi()
		},
	}
}

func ExecuteApi() {
	config := config.Initialize()

	app := fiber.New(fiber.Config{
		AppName: "Order Service API",
	})

	dependencies := routes.DependenciesInjector{
		ProductRepository: &db.ProductRepositoryDatabase{},
	}

	routes.Initialize(app, dependencies)

	app.Listen(":" + config.AppPort)
}
