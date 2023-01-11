package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/odh-data-monitor/daemon"
	"github.com/buonotti/odh-data-monitor/errors"
)

var daemonStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the daemon",
	Long: `This command starts the daemon. If the --bg flag is provided the daemon is started as a background process. In any 
case if there is already a daemon running the new one won't start.`,
	Run: func(cmd *cobra.Command, args []string) {
		bg, err := cmd.Flags().GetBool("bg")
		errors.HandleError(err)
		errors.HandleError(daemon.Setup())
		_, err = daemon.Start(bg)
		errors.HandleError(err)
	},
}

func init() {
	daemonStartCmd.Flags().Bool("bg", false, "Run the daemon in the background")
	daemonCmd.AddCommand(daemonStartCmd)
}
