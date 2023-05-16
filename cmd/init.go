package cmd

import (
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize apisense directories",
	Long:  `This command initialize apisense directories. It creates the config directory and the reports and definitions directories.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
	PostRun: func(_ *cobra.Command, _ []string) {
		log.CliLogger.Info("apisense initialized")
		log.CliLogger.WithField("directory", directories.ConfigDirectory()).Info("config directory")
		log.CliLogger.WithField("directory", directories.ReportsDirectory()).Info("reports directory")
		log.CliLogger.WithField("directory", directories.DefinitionsDirectory()).Info("definitions directory")
		log.CliLogger.WithField("directory", directories.DaemonDirectory()).Info("daemon directory")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
