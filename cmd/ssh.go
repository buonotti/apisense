package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/ssh"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Start the tui as an ssh service",
	Long:  "Start the tui as an ssh service", // TODO add desc
	Run: func(cmd *cobra.Command, args []string) {
		errors.HandleError(ssh.Start())
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
}
