package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/errors"
)

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a configuration value",
	Long:  `Get a configuration value`, // TODO: Add more info
	Run: func(cmd *cobra.Command, args []string) {
		key := cmd.Flag("key").Value.String()
		if key == "" {
			for _, key := range viper.AllKeys() {
				printConfigValue(key)
			}
		} else {
			printConfigValue(key)
		}
	},
}

func printConfigValue(key string) {
	val := viper.Get(key)
	if val == "" {
		val = "<empty>"
	}
	fmt.Printf("%s: %v\n", key, viper.Get(key)) // TODO nice formatting
}

func init() {
	configGetCmd.Flags().StringP("key", "k", "", "The key to get")

	err := configGetCmd.RegisterFlagCompletionFunc("key", validConfigKeysFunc())
	errors.HandleError(errors.SafeWrap(errors.CannotRegisterCompletionFunction, err, "cannot register completion function for config get"))

	configCmd.AddCommand(configGetCmd)
}
