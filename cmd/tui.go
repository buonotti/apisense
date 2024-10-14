package cmd

import (
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/tui"
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
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
