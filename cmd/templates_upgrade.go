package cmd

import (
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/validation/validators/pkg"
	"github.com/spf13/cobra"
)

var templatesUpgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrades one or all templates",
	Long:  "Upgrade one or all locally cached templates by pulling changes from the remote. Only works if the commit field in the lockfile is set to '*'",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) == 0 {
			err = pkg.UpgradeAll()
		} else {
			err = pkg.Upgrade(args[0])
		}
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
	},
}

func init() {
	templatesCmd.AddCommand(templatesUpgradeCmd)
}
