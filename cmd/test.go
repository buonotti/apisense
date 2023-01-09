package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/buonotti/odh-data-monitor/daemon"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			fmt.Println(daemon.Pid())
			time.Sleep(100 * time.Millisecond)
		}
	},
}

func init() {
	//rootCmd.AddCommand(testCmd)
}
