package main

import (
	"github.com/bruandreo/order-service/database"
	"github.com/bruandreo/order-service/internal/config"
	"github.com/bruandreo/order-service/internal/repositories/db"
	"github.com/bruandreo/order-service/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${method} ${path} ${status} ${latency}",
	}))

	dbConn := database.Connection(config)

	dependencies := routes.DependenciesInjector{
		ProductRepository: &db.ProductRepositoryDatabase{DB: dbConn},
		DB:                dbConn,
	}

	routes.Initialize(app, dependencies)

	app.Listen(":" + config.AppPort)
}
