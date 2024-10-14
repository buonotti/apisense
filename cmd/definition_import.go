package cmd

import (
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/imports"
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
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		content, err := os.ReadFile(file)
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
		var importer imports.Importer
		specVer, err := cmd.Flags().GetString("fmt")
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
		if specVer == "swagger2" {
			importer = &swagger.V2Importer{}
		} else if specVer == "swagger3" {
			importer = &swagger.V3Importer{}
		} else {
			log.DefaultLogger().Fatal("Invalid spec format", "format", specVer)
		}
		definitions, err := importer.Import(file, content)
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}

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
	definitionImportCmd.Flags().String("fmt", "swagger2", "Set the swagger document version")
	err := definitionImportCmd.RegisterFlagCompletionFunc("fmt", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"swagger2", "swagger3"}, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		log.DefaultLogger().Fatal(errors.CannotRegisterCompletionFunction.Wrap(err, "cannot register fmt completion func"))
	}
	err = definitionImportCmd.MarkFlagRequired("fmt")
	if err != nil {
		log.DefaultLogger().Fatal(err)
	}
	definitionCmd.AddCommand(definitionImportCmd)
}
