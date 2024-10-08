package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/theme"
	"github.com/buonotti/apisense/validation/definitions"
)

var definitionListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List definitions",
	Long:    `List definitions`, // TODO: Add more info
	Run: func(cmd *cobra.Command, _ []string) {
		definitions, err := definitions.Endpoints()
		errors.CheckErr(err)
		concise := cmd.Flags().GetBool("concise")

		if !concise {
			fmt.Println(yellowStyle().Bold(true).Render("# Definitions \n"))
		}

		for _, def := range definitions {
			printDefinition(def, concise)
		}
	},
}

func printDefinitionVerbose(definition definitions.Endpoint) {
	keyStyle := lipgloss.NewStyle().Foreground(theme.Ansi2Color(termenv.ANSIBlue)).Bold(true)
	fmt.Printf("%s: %s\n", keyStyle.Render("Filename"), definition.Name)
	fmt.Printf("%s: ", keyStyle.Render("Enabled"))
	if definition.IsEnabled {
		fmt.Printf("%v\n", greenStyle().Render("true"))
	} else {
		fmt.Printf("%v\n", redStyle().Render("false"))
	}
	fmt.Printf("%s: %s\n", keyStyle.Render("Full path"), filepath.FromSlash(definition.FullPath))
	fmt.Println()
}

func printDefinition(definition definitions.Endpoint, concise bool) {
	if !concise {
		printDefinitionVerbose(definition)
	} else {
		enabled := "enabled"
		if !definition.IsEnabled {
			enabled = "disabled"
		}
		fmt.Printf("%s (%s/%s) %s\n", definition.Name, filepath.FromSlash(directories.DefinitionsDirectory()), definition.FileName, enabled)
	}
}

func init() {
	definitionListCmd.Flags().BoolP("concise", "c", false, "Print less information")
	definitionCmd.AddCommand(definitionListCmd)
}
