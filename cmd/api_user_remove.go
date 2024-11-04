package cmd

import (
	"github.com/buonotti/apisense/v2/api/db"
	"github.com/buonotti/apisense/v2/log"
	"github.com/spf13/cobra"
)

var apiUserRemoveCmd = &cobra.Command{
	Use:               "remove [USERNAME]",
	Short:             "Remove a user",
	Long:              `This command allows to remove a user of the API.`,
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: validUsersFunc(),
	Run: func(_ *cobra.Command, args []string) {
		username := args[0]

		err := db.DeleteUser(username)
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}

		log.DefaultLogger().Info("User removed", "username", username)
	},
}

func init() {
	apiUserCmd.AddCommand(apiUserRemoveCmd)
}
