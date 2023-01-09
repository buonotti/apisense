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
		bg, err := cmd.Flags().GetBool("bg")
		errors.HandleError(err)
		errors.HandleError(daemon.Setup())
		errors.HandleError(daemon.Start(bg))
	},
}

func init() {
	daemonStartCmd.Flags().Bool("bg", false, "Run the daemon in the background")
	daemonCmd.AddCommand(daemonStartCmd)
}
