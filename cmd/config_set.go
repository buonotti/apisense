package cmd

import (
	"fmt"

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
			cobra.CheckErr(errors.CannotGetFlagValueError.New("cannot get value of flag: key"))
		}
		value := cmd.Flag("value").Value.String()
		if value == "" {
			cobra.CheckErr(errors.CannotGetFlagValueError.New("cannot get value of flag: value"))
		}

		viper.Set(key, value)
		err := viper.WriteConfig()
		cobra.CheckErr(errors.SafeWrap(errors.CannotWriteConfigError, err, "cannot write to config file"))
		fmt.Println("Config Updated!")
		printConfigValue(key, len(key))
	},
}

func init() {
	configSetCmd.Flags().StringP("key", "k", "", "The key to set")
	configSetCmd.Flags().StringP("value", "v", "", "The value to set")

	errors.CheckErr(configSetCmd.MarkFlagRequired("key"))
	errors.CheckErr(configSetCmd.MarkFlagRequired("value"))

	err := configSetCmd.RegisterFlagCompletionFunc("key", validConfigKeysFunc())
	errors.CheckErr(errors.SafeWrap(errors.CannotRegisterCompletionFunction, err, "cannot register completion function for config set"))
	configCmd.AddCommand(configSetCmd)
}
