package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/errors"
)

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a configuration value",
	Long:  `Set a configuration value`, // TODO: Add more info
	Run: func(cmd *cobra.Command, args []string) {
		key := cmd.Flag("key").Value.String()
		if key == "" {
			errors.HandleError(errors.CannotGetFlagValueError.New("cannot get value of flag: key"))
		}
		value := cmd.Flag("value").Value.String()
		if value == "" {
			errors.HandleError(errors.CannotGetFlagValueError.New("cannot get value of flag: value"))
		}

		viper.Set(key, value)
		err := viper.WriteConfig()
		errors.HandleError(errors.SafeWrap(errors.CannotWriteConfigError, err, "cannot write to config file"))
	},
}

func init() {
	configSetCmd.Flags().StringP("key", "k", "", "The key to set")
	configSetCmd.Flags().StringP("value", "v", "", "The value to set")

	errors.HandleError(configSetCmd.MarkFlagRequired("key"))
	errors.HandleError(configSetCmd.MarkFlagRequired("value"))

	err := configSetCmd.RegisterFlagCompletionFunc("key", validConfigKeysFunc())
	errors.HandleError(errors.SafeWrap(errors.CannotRegisterCompletionFunction, err, "cannot register completion function for config set"))
	configCmd.AddCommand(configSetCmd)
}
