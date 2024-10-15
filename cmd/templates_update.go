package cmd

import (
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/validation/validators/repo"
	"github.com/spf13/cobra"
)

var templatesUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update validator templates",
	Long:  "Update validator templates",
	Run: func(cmd *cobra.Command, args []string) {
		err := repo.Update()
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
		log.DefaultLogger().Info("Updated template repositories")
	},
}

func init() {
	templatesCmd.AddCommand(templatesUpdateCmd)
}
