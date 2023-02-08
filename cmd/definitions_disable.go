package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/validation"
)

var definitionsDisableCmd = &cobra.Command{
	Use:   "disable [DEFINITION]",
	Short: "Disable a definition",
	Long:  `Disable a definition`, // TODO: Add more info
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		fullPath := validation.DefinitionsLocation() + "/" + fileName
		if _, err := os.Stat(fullPath); err == nil {
			errors.CheckErr(os.Rename(fullPath, validation.DefinitionsLocation()+"/"+viper.GetString("daemon.ignore-prefix")+fileName))
			fmt.Println("Definition disabled: " + fullPath)
		} else {
			fmt.Println("Definition not found: " + fullPath)
		}
	},
	ValidArgsFunction: validEnabledDefinitionFunc(),
	Args:              cobra.ExactArgs(1),
}

func init() {
	definitionsCmd.AddCommand(definitionsDisableCmd)
}
