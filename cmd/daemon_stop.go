package cmd

import (
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/errors"
)

var daemonStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the daemon",
	Long:  `This command stops a running daemon. If there is no daemon running the command does nothing.`,
	Run: func(_ *cobra.Command, _ []string) {
		errors.CheckErr(daemon.Stop())
		log.CliLogger.Infof("Daemon stopped")
	},
}

func init() {
	daemonCmd.AddCommand(daemonStopCmd)
}
