package cmd

import (
	"fmt"
	"os"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/validation/definitions"
	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
)

var definitionCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if a definition file is valid",
	Long:  "Validates the given definition file and checks whether its structure is valid",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		content, err := os.ReadFile(file)
		if err != nil {
			log.DefaultLogger().Fatal(errors.CannotReadFileError.WrapWithNoMessage(err))
		}
		var definition definitions.Endpoint
		err = yaml.Unmarshal(content, &definition)
		if err != nil {
			log.DefaultLogger().Fatal(errors.CannotParseDefinitionFileError.WrapWithNoMessage(err))
		}
		err = definitions.ValidateDefinition(&definition)
		if err != nil {
			fmt.Printf("Definition is %s: %s\n", redStyle().Render("invalid"), err.Error())
		} else {
			fmt.Printf("Definition is %s\n", greenStyle().Render("valid"))
		}
	},
}

func init() {
	definitionCmd.AddCommand(definitionCheckCmd)
}
