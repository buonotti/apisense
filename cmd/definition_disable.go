package cmd

import (
	"os"
	"path/filepath"

	"github.com/buonotti/apisense/v2/filesystem/locations"
	"github.com/buonotti/apisense/v2/filesystem/locations/directories"
	"github.com/buonotti/apisense/v2/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var definitionDisableCmd = &cobra.Command{
	Use:               "disable [DEFINITION]",
	Short:             "Disable a definition",
	Long:              `This command is used to disable a given definition.`,
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: validEnabledDefinitionFunc(),
	Run: func(_ *cobra.Command, args []string) {
		fileName := args[0]
		fullPath := locations.DefinitionExt(fileName)
		if _, err := os.Stat(fullPath); err == nil {
			err := os.Rename(fullPath, filepath.FromSlash(directories.DefinitionsDirectory()+"/"+viper.GetString("daemon.ignore_prefix")+fileName))
			if err != nil {
				log.DefaultLogger().Fatal(err)
			}
			log.DefaultLogger().Info("Definition disabled", "filename", fileName)
		} else {
			log.DefaultLogger().Error("Definition not found", "filename", fileName)
		}
	},
}

func init() {
	definitionCmd.AddCommand(definitionDisableCmd)
}
