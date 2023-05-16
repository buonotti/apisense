package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/tui"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Start the TUI",
	Long:  `This command starts the text user interface in the current terminal. Refer to the help menu in the TUI for keybindings and more.`,
	Run: func(_ *cobra.Command, _ []string) {
		errors.CheckErr(tui.Run())
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
