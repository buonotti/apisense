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
	Long:  `This command prints "up" and pid of the daemon if there is one running or "down" and -1 as the pid if there is no daemon running.`,
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
