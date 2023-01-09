package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/odh-data-monitor/daemon"
	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/log"
)

var daemonStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of the daemon",
	Long:  `Check the status of the daemon`, // TODO add more info
	Run: func(cmd *cobra.Command, args []string) {
		status, err := daemon.Status()
		errors.HandleError(err)
		pid, err := daemon.Pid()
		errors.HandleError(err)
		log.DefaultLogger.Infof("Daemon is %s (pid %d)", status, pid)
	},
}

func init() {
	daemonCmd.AddCommand(daemonStatusCmd)
}
