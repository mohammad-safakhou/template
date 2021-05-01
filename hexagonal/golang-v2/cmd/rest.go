package cmd

import (
	"github.com/spf13/cobra"
	"template/adapters/api"
)

func init() {
	rootCmd.AddCommand(restCmd)
}

var restCmd = &cobra.Command{
	Use:   "rest-server",
	Short: "Starting rest server",
	Run: func(cmd *cobra.Command, args []string) {
		api.StartRestServer()
	},
}

