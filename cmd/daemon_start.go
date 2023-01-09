package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/odh-data-monitor/daemon"
	"github.com/buonotti/odh-data-monitor/errors"
)

var daemonStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the daemon",
	Long:  `Start the daemon`, // TODO add more info
	Run: func(cmd *cobra.Command, args []string) {
		errors.HandleError(daemon.Setup())
		errors.HandleError(daemon.Start())
	},
}

func init() {
	daemonCmd.AddCommand(daemonStartCmd)
}
