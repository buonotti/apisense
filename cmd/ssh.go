package cmd

import (
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/ssh"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Start the tui as an ssh service and the daemon",
	Long: `This command starts an ssh server that serves the tui over SSH. It also enables scp to download the reports from the server.
This command automatically starts the daemon. This behaviour can be disabled by supplying the --no-daemon flag.`,
	Run: func(_ *cobra.Command, _ []string) {
		err := ssh.Start()
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
}
