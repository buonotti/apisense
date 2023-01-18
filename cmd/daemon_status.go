package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/errors"
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
		fmt.Printf("Daemon is %s (pid %d)\n", status, pid)
	},
}

func init() {
	daemonCmd.AddCommand(daemonStatusCmd)
}
