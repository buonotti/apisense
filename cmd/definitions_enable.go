package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/validation"
)

var definitionsEnableCmd = &cobra.Command{
	Use:   "enable [DEFINITION]",
	Short: "Enable a definition",
	Long:  `Enable a definition`, // TODO: Add more info
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		fullPath := validation.DefinitionsLocation() + "/" + viper.GetString("daemon.ignore-prefix") + fileName
		if _, err := os.Stat(fullPath); err == nil {
			errors.CheckErr(os.Rename(fullPath, validation.DefinitionsLocation()+"/"+strings.TrimPrefix(fileName, viper.GetString("daemon.ignore-prefix"))))
			fmt.Println("Definition enabled: " + fullPath)
		} else {
			fmt.Println("Definition not found: " + fullPath)
		}
	},
	ValidArgsFunction: validDisabledDefinitionFunc(),
	Args:              cobra.ExactArgs(1),
}

func init() {
	definitionsCmd.AddCommand(definitionsEnableCmd)
}
