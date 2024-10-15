package cmd

import (
	"github.com/buonotti/apisense/daemon"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:     "daemon",
	Aliases: []string{"d"},
	Short:   "Manage the daemon",
	Long:    `This command is used to manage the daemon functionalities. It provides subcommands to start, stop and check the status of the daemon.`,
	Run: func(cmd *cobra.Command, _ []string) {
		cobra.CheckErr(cmd.Help())
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmd.Root().PersistentPreRun(cmd, args)
		cobra.CheckErr(daemon.Setup())
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
