package cmd

import (
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/validation/validators/pkg"
	"github.com/spf13/cobra"
)

var templatesAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a custom repo for a template",
	Long:  "Add a git repo containing the template code for an apisense validator",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		lang, err := cmd.Flags().GetString("lang")
		if err != nil {
			log.DefaultLogger().Fatal(errors.CannotGetFlagValueError.WrapWithNoMessage(err))
		}
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			log.DefaultLogger().Fatal(errors.CannotGetFlagValueError.WrapWithNoMessage(err))
		}
		branch, _ := cmd.Flags().GetString("branch")
		commit, _ := cmd.Flags().GetString("commit")
		err = pkg.AddTemplateSource(lang, url, branch, commit)
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
		log.DefaultLogger().Info("Added template for language. Run 'apisense templates update' to fetch the new template", "lang", lang)
	},
}

func init() {
	templatesAddCmd.Flags().StringP("lang", "l", "", "The language of the template")
	templatesAddCmd.Flags().StringP("url", "u", "", "The remote repo url to clone the template from")
	templatesAddCmd.Flags().StringP("branch", "b", "main", "The branch to use")
	templatesAddCmd.Flags().StringP("commit", "c", "*", "The commit to use. Set to * to use the latest")
	err := templatesAddCmd.MarkFlagRequired("lang")
	if err != nil {
		log.DefaultLogger().Fatal(errors.CannotMarkFlagRequiredError.WrapWithNoMessage(err))
	}

	err = templatesAddCmd.MarkFlagRequired("url")
	if err != nil {
		log.DefaultLogger().Fatal(errors.CannotMarkFlagRequiredError.WrapWithNoMessage(err))
	}

	templatesCmd.AddCommand(templatesAddCmd)
}
