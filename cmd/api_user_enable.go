package cmd

import (
	"github.com/buonotti/apisense/api/db"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
)

var apiUserEnableCmd = &cobra.Command{
	Use:   "enable [USERNAME]",
	Short: "Enable a user",
	Long:  `This command allows to enable a user of the API.`,
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		username := args[0]

		err := db.EnableUser(username)
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}

		log.DefaultLogger().Info("User enabled", "username")
	},
	ValidArgsFunction: validDisabledUserFunc(),
}

func init() {
	apiUserCmd.AddCommand(apiUserEnableCmd)
}
