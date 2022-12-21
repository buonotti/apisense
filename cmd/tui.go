package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/tui"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Start the TUI",
	Long:  `Start the TUI`, // TODO
	Run: func(cmd *cobra.Command, args []string) {
		errors.HandleError(tui.Run())
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
