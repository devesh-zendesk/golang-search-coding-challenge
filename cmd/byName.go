package cmd

import (
	"github.com/spf13/cobra"
)

var search_by_name_cmd = &cobra.Command{
	Use:   "by_name",
	Short: "Search Users details by name",
	Long:  `Search Users details by name.`,
	Run: func(cmd *cobra.Command, args []string) {
		//to be implemented
	},
}

func init() {
	rootCmd.AddCommand(search_by_name_cmd)
}
