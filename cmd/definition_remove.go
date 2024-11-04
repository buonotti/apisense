package cmd

import (
	"os"

	"github.com/buonotti/apisense/v2/filesystem/locations"
	"github.com/buonotti/apisense/v2/log"
	"github.com/spf13/cobra"
)

var definitionRemoveCmd = &cobra.Command{
	Use:               "remove [DEFINITIONS]...",
	Short:             "Removes one or more definitions",
	Long:              "Removes the definition files associated to the given definitions.",
	Args:              cobra.MinimumNArgs(1),
	ValidArgsFunction: validDefinitionsFunc(),
	Run: func(_ *cobra.Command, args []string) {
		for _, fileName := range args {
			fullPath := locations.DefinitionExt(fileName)
			err := os.Remove(fullPath)
			if err != nil {
				log.DefaultLogger().Fatal(err)
			}
			log.DefaultLogger().Info("Definition removed", "filename", fullPath)
		}
	},
}

func init() {
	definitionCmd.AddCommand(definitionRemoveCmd)
}
