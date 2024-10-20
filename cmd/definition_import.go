package cmd

import (
	"fmt"
	"os"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/imports"
	"github.com/buonotti/apisense/imports/swagger"
	"github.com/buonotti/apisense/log"
	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
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
			log.DefaultLogger().Fatal(errors.CannotReadFileError.WrapWithNoMessage(err))
		}
		var importer imports.Importer
		specVer, err := cmd.Flags().GetString("spec")
		if err != nil {
			log.DefaultLogger().Fatal(errors.CannotGetFlagValueError.WrapWithNoMessage(err))
		}

		if specVer == "swagger2" {
			importer = &swagger.V2Importer{}
		} else if specVer == "swagger3" {
			importer = &swagger.V3Importer{}
		} else {
			log.DefaultLogger().Fatal("Invalid spec format", "format", specVer)
		}
		definitions, err := importer.Import(content)
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}

		for _, def := range definitions {
			defStr, err := yaml.Marshal(def)
			if err != nil {
				log.DefaultLogger().Fatal(errors.CannotSerializeItemError.WrapWithNoMessage(err))
			}
			if write, _ := cmd.Flags().GetBool("write"); write {
				err = os.WriteFile(def.FileName, defStr, os.ModePerm)
				if err != nil {
					log.DefaultLogger().Fatal(errors.CannotWriteFileError.WrapWithNoMessage(err))
				}
				log.DefaultLogger().Info("Definition file successfully created", "filename", def.FullPath)
			} else {
				fmt.Printf("\n--- %s\n%s\n", def.FullPath, defStr)
			}
		}

		log.DefaultLogger().Info("Successfully imported definitions", "amount", len(definitions))
	},
}

func init() {
	definitionImportCmd.Flags().String("spec", "swagger2", "Set the swagger document version")
	err := definitionImportCmd.RegisterFlagCompletionFunc("spec", validImportFormatsFunc())
	if err != nil {
		log.DefaultLogger().Fatal(errors.CannotRegisterCompletionFunction.WrapWithNoMessage(err))
	}
	err = definitionImportCmd.MarkFlagRequired("spec")
	if err != nil {
		log.DefaultLogger().Fatal(errors.CannotMarkFlagRequiredError.WrapWithNoMessage(err))
	}

	definitionImportCmd.Flags().BoolP("write", "w", false, "Also write the definitions to disk")
	definitionCmd.AddCommand(definitionImportCmd)
}
