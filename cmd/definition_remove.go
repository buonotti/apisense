package cmd

import (
	"os"
	"path/filepath"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
)

var definitionRemoveCmd = &cobra.Command{
	Use:               "remove [DEFINITIONS]...",
	Short:             "Removes one or more definitions",
	Long:              "Removes the definition files associated to the given definitions.",
	Args:              cobra.MinimumNArgs(1),
	ValidArgsFunction: validDefinitionsFunc(),
	Run: func(_ *cobra.Command, args []string) {
		for _, definition := range args {
			fullPath := filepath.FromSlash(directories.DefinitionsDirectory() + "/" + definition)
			err := os.Remove(fullPath)
			errors.CheckErr(err)
			log.DefaultLogger().Info("Definition removed", "filename", fullPath)
		}
	},
}

func init() {
	definitionCmd.AddCommand(definitionRemoveCmd)
}
