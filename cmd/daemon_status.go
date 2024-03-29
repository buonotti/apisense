package cmd

import (
	"fmt"

	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/errors"
)

var daemonStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of the daemon",
	Long:  `This command prints "up" and pid of the daemon if there is one running or "down" and -1 as the pid if there is no daemon running.`,
	Run: func(_ *cobra.Command, _ []string) {
		status, err := daemon.Status()
		errors.CheckErr(err)
		pid, err := daemon.Pid()
		errors.CheckErr(err)
		var styledStatus string
		if status == daemon.UpStatus {
			styledStatus = greenStyle().Bold(true).Render(string(status))
		} else {
			styledStatus = redStyle().Bold(true).Render(string(status))
		}
		var styledPid string
		if pid == -1 {
			styledPid = redStyle().Italic(true).Render(fmt.Sprintf("%d", pid))
		} else {
			styledPid = greenStyle().Italic(true).Render(fmt.Sprintf("%d", pid))
		}
		log.CliLogger.WithField("status", styledStatus).WithField("pid", styledPid).Info("daemon status retrieved")
	},
}

func init() {
	daemonCmd.AddCommand(daemonStatusCmd)
}
