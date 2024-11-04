package cmd

import (
	"github.com/buonotti/apisense/v2/daemon"
	"github.com/buonotti/apisense/v2/log"
	"github.com/buonotti/apisense/v2/tui"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Start the TUI",
	Long:  `This command starts the text user interface in the current terminal. Refer to the help menu in the TUI for keybindings and more.`,
	Args:  cobra.NoArgs,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(daemon.Setup())
	},
	Run: func(_ *cobra.Command, _ []string) {
		err := tui.Run()
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
