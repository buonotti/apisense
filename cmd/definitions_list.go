package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/validation"
)

var definitionsListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List definitions",
	Long:    `List definitions`, // TODO: Add more info
	Run: func(cmd *cobra.Command, args []string) {
		definitions, err := validation.EndpointDefinitions()
		errors.HandleError(err)

		for _, def := range definitions {
			fmt.Printf("%s (%s/%s)\n", def.Name, validation.DefinitionsLocation(), def.FileName)
		}
	},
}

func init() {
	definitionsCmd.AddCommand(definitionsListCmd)
}
