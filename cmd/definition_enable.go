package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/buonotti/apisense/filesystem/locations"
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var definitionEnableCmd = &cobra.Command{
	Use:               "enable [DEFINITION]",
	Short:             "Enable a definition",
	Long:              `Enable a definition`,
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: validDisabledDefinitionFunc(),
	Run: func(_ *cobra.Command, args []string) {
		fileName := args[0]
		fullPath := locations.DefinitionExt(fileName)
		if _, err := os.Stat(fullPath); err == nil {
			err = os.Rename(fullPath, filepath.FromSlash(directories.DefinitionsDirectory()+"/"+strings.TrimPrefix(fileName, viper.GetString("daemon.ignore_prefix"))))
			if err != nil {
				log.DefaultLogger().Fatal(err)
			}
			log.DefaultLogger().Info("Definition enabled", "filename", fileName)
		} else {
			log.DefaultLogger().Error("definition not found", "filename", fileName)
		}
	},
}

func init() {
	definitionCmd.AddCommand(definitionEnableCmd)
}
