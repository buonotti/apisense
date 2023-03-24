package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/errors"
)

var daemonStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the daemon",
	Long:  `This command starts the daemon. In any case if there is already a daemon running the new one won't start.`,
	Run: func(cmd *cobra.Command, args []string) {
		force, err := cmd.Flags().GetBool("force")
		if err != nil {
			errors.CheckErr(errors.CannotGetFlagValueError.Wrap(err, "cannot get value of flag: force"))
		}
		err = daemon.Start(force)
		errors.CheckErr(err)
	},
}

func init() {
	daemonStartCmd.Flags().BoolP("force", "f", false, "Force validation upon startup")
	daemonCmd.AddCommand(daemonStartCmd)
}
