package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/odh-data-monitor/daemon"
	"github.com/buonotti/odh-data-monitor/errors"
)

var daemonStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the daemon",
	Long:  `Start the daemon`, // TODO
	Run: func(cmd *cobra.Command, args []string) {
		errors.HandleError(daemon.Start())
	},
}

func init() {
	daemonCmd.AddCommand(daemonStartCmd)
}
