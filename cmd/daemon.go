package cmd

import (
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Manage the daemon",
	Long:  `This command is used to manage the daemon functionalities. It provides subcommands to start, stop and check the status of the daemon.`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(cmd.Help())
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
