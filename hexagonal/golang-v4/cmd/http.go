package cmd

import (
	"git.siz-tel.com/charging/template/internal/gateway/rest"
	"git.siz-tel.com/charging/template/logger"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(httpCmd)
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "launching the http rest listen server",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Logger.Info("http rest server is starting")
		rest.Start()
	},
}
