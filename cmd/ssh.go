package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/ssh"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Start the tui as an ssh service and the daemon",
	Long: `This command starts an ssh server that serves the tui over SSH. It also enables scp to download the reports from the server.
This command automatically starts the daemon. This behaviour can be disabled by supplying the --no-daemon flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		errors.HandleError(ssh.Start())
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
}
