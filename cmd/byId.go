package cmd

import (
	"github.com/spf13/cobra"
)

var search_by_id_cmd = &cobra.Command{
	Use:   "by_id",
	Short: "Search Users details by ID",
	Long:  `Search Users details by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		//to be implemented
	},
}

func init() {
	rootCmd.AddCommand(search_by_id_cmd)
}
