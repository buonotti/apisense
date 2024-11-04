package cmd

import (
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/validation/validators"
	"github.com/spf13/cobra"
)

var validatorsDiscoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "Tries to discover and add the binary of a validator to the config",
	Long:  "Uses project files to determine the language of the validator the adds the default produced variable to the main config",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		override, _ := cmd.Flags().GetBool("override")
		err := validators.AutoDiscoverExternal(override)
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
	},
}

func init() {
	validatorsDiscoverCmd.Flags().Bool("override", false, "Override existing validators with the same name")

	validatorsCmd.AddCommand(validatorsDiscoverCmd)
}
