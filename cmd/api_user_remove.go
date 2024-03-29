package cmd

import (
	"github.com/buonotti/apisense/api/db"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
)

var apiUserRemoveCmd = &cobra.Command{
	Use:   "remove [USERNAME]",
	Short: "Remove a user",
	Long:  `This command allows to remove a user of the API.`,
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		username := args[0]

		err := db.DeleteUser(username)
		errors.CheckErr(err)

		log.CliLogger.WithField("username", username).Info("user removed")
	},
	ValidArgsFunction: validUsersFunc(),
}

func init() {
	apiUserCmd.AddCommand(apiUserRemoveCmd)
}
