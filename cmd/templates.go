package cmd

import (
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/validation/validators/repo"
	"github.com/spf13/cobra"
)

var templatesCmd = &cobra.Command{
	Use:   "templates",
	Short: "Manage validator templates",
	Long:  "This command is used to install and update validator templates from official and non-official repositories",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := repo.Update()
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(templatesCmd)
}
