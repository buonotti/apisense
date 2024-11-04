package cmd

import (
	"fmt"

	"github.com/buonotti/apisense/v2/filesystem/locations/directories"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Print apisense env",
	Long:  "Prints out all the relevant apisense directories/files",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s: %s\n", "Config dir", directories.ConfigDirectory())
		fmt.Printf("%s: %s\n", "Definitions dir", directories.DefinitionsDirectory())
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
}
