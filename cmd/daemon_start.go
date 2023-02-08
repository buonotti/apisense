package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/errors"
)

var daemonStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the daemon",
	Long: `This command starts the daemon. If the --bg flag is provided the daemon is started as a background process. In any 
case if there is already a daemon running the new one won't start.`,
	Run: func(cmd *cobra.Command, args []string) {
		bg, err := cmd.Flags().GetBool("background")
		errors.CheckErr(errors.SafeWrap(errors.CannotGetFlagValueError, err, "cannot get value of flag: background"))
		force, err := cmd.Flags().GetBool("force")
		errors.CheckErr(errors.SafeWrap(errors.CannotGetFlagValueError, err, "cannot get value of flag: force"))
		_, err = daemon.Start(bg, force)
		errors.CheckErr(err)
		if bg {
			fmt.Printf("Daemon now running in background\n")
		}
	},
}

func init() {
	daemonStartCmd.Flags().BoolP("force", "f", false, "Force validation upon startup")
	daemonStartCmd.Flags().Bool("background", false, "Run the daemon in the background")
	daemonCmd.AddCommand(daemonStartCmd)
}
