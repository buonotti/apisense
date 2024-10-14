package cmd

import (
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/validation/validators/repo"
	"github.com/spf13/cobra"
)

var templatesCreateCmd = &cobra.Command{
	Use:   "create [NAME]",
	Short: "Create a new validator from a template",
	Long:  "Create a new external apisense validator from an installed validator template. See 'apisense templates list' for a list of all available validators",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lang, err := cmd.Flags().GetString("lang")
		if err != nil {
			log.DefaultLogger().Fatal(errors.CannotGetFlagValueError.Wrap(err, "cannot get value of flag 'lang'"))
		}
		dest := args[0]
		err = repo.Create(lang, dest)
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
	},
}

func init() {
	templatesCreateCmd.Flags().StringP("lang", "l", "", "The language of the template to use")

	templatesCmd.AddCommand(templatesCreateCmd)
}
