package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "search",
		Short: "Searching CLI Application",
		Long:  `Searching CLI Application.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
