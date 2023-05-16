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
	Long:  `Disable a definition`, // TODO: Add more info
	Run: func(_ *cobra.Command, args []string) {
		fileName := args[0]
		fullPath := filepath.FromSlash(directories.DefinitionsDirectory() + "/" + fileName)
		if _, err := os.Stat(fullPath); err == nil {
			errors.CheckErr(os.Rename(fullPath, filepath.FromSlash(directories.DefinitionsDirectory()+"/"+viper.GetString("daemon.ignore_prefix")+fileName)))
			log.CliLogger.Infof("Definition disabled: %s", fullPath)
		} else {
			log.CliLogger.Infof("Definition not found: %s", fullPath)
		}
	},
	ValidArgsFunction: validEnabledDefinitionFunc(),
	Args:              cobra.ExactArgs(1),
}

func init() {
	definitionCmd.AddCommand(definitionDisableCmd)
}
