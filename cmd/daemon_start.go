package cmd

import (
	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var daemonStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the daemon",
	Long:  `This command starts the daemon. In any case if there is already a daemon running the new one won't start.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		force, err := cmd.Flags().GetBool("force")
		force = force || viper.GetBool("daemon.run_at_startup")
		if err != nil {
			log.DefaultLogger().Fatal(errors.CannotGetFlagValueError.Wrap(err, "cannot get value of flag: force"))
		}
		err = daemon.Start(force)
		if err != nil {
			log.DaemonLogger().Fatal(err)
		}
	},
}

func init() {
	daemonStartCmd.Flags().BoolP("force", "f", false, "Force validation upon startup")
	daemonCmd.AddCommand(daemonStartCmd)
}
