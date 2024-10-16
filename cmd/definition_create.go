package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	_ "embed"

	"github.com/buonotti/apisense/examples"
	"github.com/buonotti/apisense/filesystem/locations/directories"
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
		fileName := filepath.FromSlash(directories.DefinitionsDirectory() + "/" + name + ".apisensedef.yml")
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
