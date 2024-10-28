package cmd

import (
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/validation/validators/pkg"
	"github.com/spf13/cobra"
)

var templatesAddrepoCmd = &cobra.Command{
	Use:   "add-repo",
	Short: "Add a remote template repo",
	Long:  "Adds a remote repository containing git repos that are validator templates",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.DefaultLogger().Fatal(errors.CannotGetFlagValueError.WrapWithNoMessage(err))
		}

		url, err := cmd.Flags().GetString("url")
		if err != nil {
			log.DefaultLogger().Fatal(errors.CannotGetFlagValueError.WrapWithNoMessage(err))
		}

		err = pkg.AddRepo(name, url, pkg.GitHub)
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
	},
}

func init() {
	templatesAddrepoCmd.Flags().StringP("name", "n", "", "The name of the repo")
	templatesAddrepoCmd.Flags().StringP("url", "u", "", "The url of the repo")

	err := templatesAddrepoCmd.MarkFlagRequired("name")
	if err != nil {
		log.DefaultLogger().Fatal(errors.CannotMarkFlagRequiredError.WrapWithNoMessage(err))
	}
	err = templatesAddrepoCmd.MarkFlagRequired("url")
	if err != nil {
		log.DefaultLogger().Fatal(errors.CannotMarkFlagRequiredError.WrapWithNoMessage(err))
	}

	templatesCmd.AddCommand(templatesAddrepoCmd)
}
