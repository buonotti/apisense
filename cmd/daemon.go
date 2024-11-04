package cmd

import (
	"github.com/buonotti/apisense/v2/daemon"
	"github.com/buonotti/apisense/v2/log"
	clog "github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:     "daemon",
	Aliases: []string{"d"},
	Short:   "Manage the daemon",
	Long:    `This command is used to manage the daemon functionalities. It provides subcommands to start, stop and check the status of the daemon.`,
	Args:    cobra.NoArgs,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmd.Root().PersistentPreRun(cmd, args)
		cobra.CheckErr(daemon.Setup())
		if ll, err := cmd.Root().PersistentFlags().GetString("log-level"); err == nil && ll != "" {
			parsed, err := clog.ParseLevel(ll)
			if err == nil {
				clog.SetLevel(parsed)
				clog.SetReportCaller(clog.GetLevel() == clog.DebugLevel)
			} else {
				log.DefaultLogger().Warn("Log level invalid. Falling back to config", "reason", err.Error())
			}
		}
	},
	Run: func(cmd *cobra.Command, _ []string) {
		cobra.CheckErr(cmd.Help())
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
