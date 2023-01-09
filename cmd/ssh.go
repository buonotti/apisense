package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/ssh"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Start the tui as an ssh service and the daemon",
	Long:  "Start the tui as an ssh service and the daemon", // TODO add desc
	Run: func(cmd *cobra.Command, args []string) {
		noDaemon, err := cmd.Flags().GetBool("no-daemon")
		errors.HandleError(err)
		errors.HandleError(ssh.Start(!noDaemon))
	},
}

func init() {
	sshCmd.Flags().Bool("no-daemon", false, "Do not start the daemon")
	rootCmd.AddCommand(sshCmd)
}
