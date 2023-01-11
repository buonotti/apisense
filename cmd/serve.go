package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/odh-data-monitor/api"
	"github.com/buonotti/odh-data-monitor/daemon"
	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/ssh"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the app by starting all the services",
	Long:  `This command starts all the services needed to serve the app. It starts the daemon, the ssh server and the tui.`,
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			cmd, err := daemon.Start(true)
			defer func() {
				if cmd != nil {
					errors.HandleError(cmd.Wait())
				}
			}()
			errors.HandleError(err)
		}()
		go func() {
			errors.HandleError(ssh.Start())
		}()
		go func() {
			errors.HandleError(api.Start())
		}()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
