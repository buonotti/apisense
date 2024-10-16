package cmd

import (
	"fmt"

	"github.com/buonotti/apisense/api/db"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
)

var apiUserListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List users",
	Aliases: []string{"ls"},
	Args:    cobra.NoArgs,
	Long:    `This command allows to list the users of the API.`,
	Run: func(_ *cobra.Command, _ []string) {
		users, err := db.ListUsers()
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}

		for _, user := range users {
			if user.Enabled {
				fmt.Printf("%s %s\n", user.Username, greenStyle().Render("enabled"))
			} else {
				fmt.Printf("%s %s\n", user.Username, redStyle().Render("disabled"))
			}
		}
	},
}

func init() {
	apiUserCmd.AddCommand(apiUserListCmd)
}
