package cmd

import (
	"os"
	"path/filepath"

	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/errors"
)

var definitionDisableCmd = &cobra.Command{
	Use:   "disable [DEFINITION]",
	Short: "Disable a definition",
	Long:  `This command is used to disable a given definition.`,
	Run: func(_ *cobra.Command, args []string) {
		fileName := args[0]
		fullPath := filepath.FromSlash(directories.DefinitionsDirectory() + "/" + fileName)
		if _, err := os.Stat(fullPath); err == nil {
			errors.CheckErr(os.Rename(fullPath, filepath.FromSlash(directories.DefinitionsDirectory()+"/"+viper.GetString("daemon.ignore_prefix")+fileName)))
			log.DefaultLogger().Info("Definition disabled", "filename", fileName)
		} else {
			log.DefaultLogger().Error("Definition not found", "filename", fileName)
		}
	},
	ValidArgsFunction: validEnabledDefinitionFunc(),
	Args:              cobra.ExactArgs(1),
}

func init() {
	definitionCmd.AddCommand(definitionDisableCmd)
}
