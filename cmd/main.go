package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{}
	rootCmd.AddCommand(ApiCommand())
	rootCmd.AddCommand(MigrationCommand())
	rootCmd.Execute()
}
