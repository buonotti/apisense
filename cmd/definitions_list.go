package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
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
		errors.CheckErr(err)
		concise := cmd.Flag("concise").Value.String() == "true"

		if !concise {
			fmt.Println(yellowStyle().Bold(true).Render("# Definitions \n"))
		}

		for _, def := range definitions {
			printDefinition(def, concise)
		}
	},
}

func printDefinitionVerbose(definition validation.EndpointDefinition) {
	keyStyle := lipgloss.NewStyle().Bold(true)
	fmt.Printf("%s: %s\n", keyStyle.Render("Filename"), definition.Name)
	fmt.Printf("%s: %s\n", keyStyle.Render("Full path"), validation.DefinitionsLocation()+definition.Name)
	fmt.Printf("%s: %s\n", keyStyle.Render("Base url"), definition.BaseUrl)
	fmt.Println()
}

func printDefinition(definition validation.EndpointDefinition, concise bool) {
	if !concise {
		printDefinitionVerbose(definition)
	} else {
		fmt.Printf("%s (%s/%s)\n", definition.Name, validation.DefinitionsLocation(), definition.FileName)
	}
}

func init() {
	definitionsListCmd.Flags().BoolP("concise", "c", false, "Print less information")

	definitionsCmd.AddCommand(definitionsListCmd)
}
