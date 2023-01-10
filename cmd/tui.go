package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/tui"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Start the TUI",
	Long:  `This command starts the text user interface in the current terminal. Refer to the help menu in the TUI for keybindings and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		errors.HandleError(tui.Run())
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
