package cmd

import (
	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
)

var daemonStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the daemon",
	Long:  `This command stops a running daemon. If there is no daemon running the command does nothing.`,
	Args:  cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		err := daemon.Stop()
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
		log.DefaultLogger().Info("Daemon stopped")
	},
}

func init() {
	daemonCmd.AddCommand(daemonStopCmd)
}
