package cmd

import (
	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/tui"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Start the TUI",
	Long:  `This command starts the text user interface in the current terminal. Refer to the help menu in the TUI for keybindings and more.`,
	Run: func(_ *cobra.Command, _ []string) {
		err := tui.Run()
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(daemon.Setup())
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
