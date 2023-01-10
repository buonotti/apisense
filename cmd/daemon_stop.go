package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/odh-data-monitor/daemon"
	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/log"
)

var daemonStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the daemon",
	Long:  `This command stops a running daemon. If there is no daemon running the command does nothing.`,
	Run: func(cmd *cobra.Command, args []string) {
		errors.HandleError(daemon.Stop())
		log.DefaultLogger.Info("Daemon stopped")
	},
}

func init() {
	daemonCmd.AddCommand(daemonStopCmd)
}
