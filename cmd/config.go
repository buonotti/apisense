package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/errors"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Long:  `Manage configuration`, // TODO: Add more info
	Run: func(cmd *cobra.Command, args []string) {
		errors.HandleError(cmd.Help())
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
