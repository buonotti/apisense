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
			log.DefaultLogger().Fatal(errors.CannotGetFlagValueError.WrapWithNoMessage(err))
		}
		force, err := cmd.Flags().GetBool("force")
		if err != nil {
			log.DefaultLogger().Fatal(errors.CannotGetFlagValueError.WrapWithNoMessage(err))
		}

		dest := args[0]
		err = repo.Create(lang, dest, force)
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
	},
}

func init() {
	templatesCreateCmd.Flags().StringP("lang", "l", "", "The language of the template to use")
	err := templatesCreateCmd.MarkFlagRequired("lang")
	if err != nil {
		log.DefaultLogger().Fatal(errors.CannotMarkFlagRequiredError.WrapWithNoMessage(err))
	}
	templatesCreateCmd.Flags().BoolP("force", "f", false, "Force the creation. Overrides already existing validators. Use with caution")

	templatesCmd.AddCommand(templatesCreateCmd)
}
