package cmd

import (
	"fmt"
	"os"

	_ "embed"

	"github.com/buonotti/apisense/examples"
	"github.com/buonotti/apisense/filesystem/locations"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
)

var definitionCreateCmd = &cobra.Command{
	Use:   "create [NAME]",
	Short: "Creates a new definition file",
	Long:  `Creates a new definition file with the needed boilerplate and the given name`,
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		name := args[0]
		fileName := locations.Definition(name)
		err := os.WriteFile(fileName, []byte(fmt.Sprintf(examples.Boilerplate, name)), os.ModePerm)
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
		log.DefaultLogger().Info("Definition file successfully created", "filename", fileName)
	},
}

func init() {
	definitionCmd.AddCommand(definitionCreateCmd)
}
