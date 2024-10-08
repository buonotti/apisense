package cmd

import (
	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
)

var daemonReloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload the daemon",
	Long:  `Reload the daemon`, // TODO: Add more info
	Run: func(_ *cobra.Command, _ []string) {
		err := daemon.ReloadDaemon()
		if err != nil {
			log.DefaultLogger().Info("Reloaded daemon pipeline items")
		} else {
			log.DefaultLogger().Error("Cannot reload daemon pipeline items")
		}
	},
}

func init() {
	daemonCmd.AddCommand(daemonReloadCmd)
}
