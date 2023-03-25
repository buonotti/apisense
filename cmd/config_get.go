package cmd

import (
	"fmt"
	"sort"

	"github.com/buonotti/apisense/log"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/util"
)

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a configuration value",
	Long:  `Get a configuration value`, // TODO: Add more info
	Run: func(cmd *cobra.Command, args []string) {
		key := cmd.Flag("key").Value.String()
		if key == "" {
			allKeys := viper.AllKeys()
			sort.Strings(allKeys)
			keyLengths := util.Map(allKeys, func(key string) int { return len(key) })
			maxKeyLength := util.Max(keyLengths)
			for _, key := range allKeys {
				printConfigValue(key, maxKeyLength)
			}
		} else {
			printConfigValue(key, len(key))
		}
	},
}

func printConfigValue(key string, maxKeyLength int) {
	val := viper.Get(key)

	key = util.Pad(key, maxKeyLength)

	styledKey := lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("%s = ", key))
	styledVal := yellowStyle().Render(fmt.Sprintf("%v", val))
	switch val := val.(type) {
	case bool:
		if val {
			styledVal = greenStyle().Render(fmt.Sprintf("%v", val))
		} else {
			styledVal = redStyle().Render(fmt.Sprintf("%v", val))
		}
	case int64:
		styledVal = blueStyle().Render(fmt.Sprintf("%v", val))
	case float64:
		styledVal = blueStyle().Render(fmt.Sprintf("%v", val))
	case string:
		if val == "" {
			val = "<empty>"
			styledVal = greyedOutStyle().Italic(true).Render(fmt.Sprintf("%v", val))
		}
	}
	log.CliLogger.Infof("%s%s", styledKey, styledVal)
}

func init() {
	configGetCmd.Flags().StringP("key", "k", "", "The key to get")

	err := configGetCmd.RegisterFlagCompletionFunc("key", validConfigKeysFunc())
	errors.CheckErr(errors.SafeWrap(errors.CannotRegisterCompletionFunction, err, "cannot register completion function for config get"))

	configCmd.AddCommand(configGetCmd)
}
