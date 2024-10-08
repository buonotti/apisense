package cmd

import (
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem"
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
		errors.CheckErr(filesystem.Setup())
		log.DefaultLogger().Info("Apisense initialized",
			"config_directory", directories.ConfigDirectory(),
			"report_directory", directories.ReportsDirectory(),
			"definitions_directory", directories.DefinitionsDirectory(),
			"deamon_directory", directories.DaemonDirectory())
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
