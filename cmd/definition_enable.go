package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/errors"
)

var definitionEnableCmd = &cobra.Command{
	Use:   "enable [DEFINITION]",
	Short: "Enable a definition",
	Long:  `Enable a definition`, // TODO: Add more info
	Run: func(_ *cobra.Command, args []string) {
		fileName := args[0]
		fullPath := filepath.FromSlash(directories.DefinitionsDirectory() + "/" + fileName)
		if _, err := os.Stat(fullPath); err == nil {
			errors.CheckErr(os.Rename(fullPath, filepath.FromSlash(directories.DefinitionsDirectory()+"/"+strings.TrimPrefix(fileName, viper.GetString("daemon.ignore_prefix")))))
			log.DefaultLogger().Info("Definition enabled", "filename", fileName)
		} else {
			log.DefaultLogger().Error("definition not found", "filename", fileName)
		}
	},
	ValidArgsFunction: validDisabledDefinitionFunc(),
	Args:              cobra.ExactArgs(1),
}

func init() {
	definitionCmd.AddCommand(definitionEnableCmd)
}
