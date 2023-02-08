package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/api"
	"github.com/buonotti/apisense/errors"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start the api server",
	Long:  `This command starts the api server.`,
	Run: func(cmd *cobra.Command, args []string) {
		errors.CheckErr(api.Start())
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
