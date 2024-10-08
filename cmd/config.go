package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Long:  `Manage configuration`, // TODO: Add more info
	Run: func(cmd *cobra.Command, _ []string) {
		cobra.CheckErr(cmd.Help())
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
