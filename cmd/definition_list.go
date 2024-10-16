package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/theme"
	"github.com/buonotti/apisense/validation/definitions"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

var definitionListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List definitions",
	Long:    `List all definitions`,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		definitions, err := definitions.Endpoints()
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
		concise := cmd.Flag("concise").Value.String() == "true"

		if len(definitions) == 0 {
			fmt.Println("No definitions")
		} else {
			if !concise {
				fmt.Println(yellowStyle().Bold(true).Render("# Definitions \n"))
			}
			for _, def := range definitions {
				printDefinition(def, concise)
			}
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
