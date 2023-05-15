package cmd

import (
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize apisense directories",
	Long:  `This command initialize apisense directories. It creates the config directory and the reports and definitions directories.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
	PostRun: func(cmd *cobra.Command, args []string) {
		log.CliLogger.Info("apisense initialized")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
