package cmd

import (
	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Manage the reports",
	Long:  "", // TODO
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(cmd.Help())
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
