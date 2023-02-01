package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/errors"
)

var definitionsCmd = &cobra.Command{
	Use:   "definitions",
	Short: "Manage definitions",
	Long:  `Manage definitions`, // TODO: Add more info
	Run: func(cmd *cobra.Command, args []string) {
		errors.CheckErr(cmd.Help())
	},
}

func init() {
	rootCmd.AddCommand(definitionsCmd)
}
