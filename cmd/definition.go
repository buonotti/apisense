package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/errors"
)

var definitionCmd = &cobra.Command{
	Use:   "definition",
	Short: "Manage definitions",
	Long:  `Manage definitions`, // TODO: Add more info
	Run: func(cmd *cobra.Command, args []string) {
		errors.CheckErr(cmd.Help())
	},
}

func init() {
	rootCmd.AddCommand(definitionCmd)
}
