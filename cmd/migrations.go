package main

import (
	"github.com/bruandreo/order-service/database"
	"github.com/bruandreo/order-service/internal/config"
	"github.com/spf13/cobra"
)

func MigrationCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "migrations",
		Short: "Execute Migration",
		Run: func(cmd *cobra.Command, args []string) {
			ExecuteMigrations()
		},
	}
}

func ExecuteMigrations() {
	config := config.Initialize()
	database.RunMigrations(config)
}
