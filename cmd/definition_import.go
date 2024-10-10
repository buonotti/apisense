package cmd

import (
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/imports/swagger"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
	"os"
)

var definitionImportCmd = &cobra.Command{
	Use:   "import [FILE]",
	Short: "Import a new definition file",
	Long:  `Import a new definition file from a supported file type`,
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		file := args[0]
		content, err := os.ReadFile(file)
		errors.CheckErr(err)
		importer := swagger.V2Importer{}
		definitions, err := importer.Import(file, content)
		errors.CheckErr(err)

		for _, def := range definitions {
			// defStr, err := yaml.Marshal(def)
			// errors.CheckErr(err)
			// err = os.WriteFile(fileName, defStr, os.ModePerm)
			// errors.CheckErr(err)
			log.DefaultLogger().Info("Definition file successfully created", "filename", def.FullPath)
		}

		log.DefaultLogger().Info("Successfully imported definitions", "amount", len(definitions))
	},
}

func init() {
	definitionCmd.AddCommand(definitionImportCmd)
}
