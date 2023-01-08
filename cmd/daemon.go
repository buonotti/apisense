package cmd

import (
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Manage the daemon",
	Long:  `Manage the daemon`, // TODO add more info
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(cmd.Help())
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
