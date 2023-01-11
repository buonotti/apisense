package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/odh-data-monitor/api"
	"github.com/buonotti/odh-data-monitor/errors"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start the api server",
	Long:  `This command starts the api server.`,
	Run: func(cmd *cobra.Command, args []string) {
		errors.HandleError(api.Start())
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
