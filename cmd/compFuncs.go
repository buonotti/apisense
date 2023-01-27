package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/conversion"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation"
)

func validReportsFunc() func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		reports, _ := validation.Reports()
		return util.Map(reports, func(r validation.Report) string {
			return fmt.Sprintf("%s\t%s with %d result(s)", r.Id, time.Time(r.Time).Format("2006-01-02 at 15-04-05.000Z"), len(r.Endpoints))
		}), cobra.ShellCompDirectiveNoFileComp
	}
}

func validConfigKeysFunc() func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		keys := viper.AllKeys()
		return keys, cobra.ShellCompDirectiveNoFileComp
	}
}

func validFormatsFunc() func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return conversion.Converters(), cobra.ShellCompDirectiveNoFileComp
	}
}

func validDefinitionsFunc() func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		definitions, _ := validation.EndpointDefinitions()
		return util.Map(definitions, func(d validation.EndpointDefinition) string {
			return fmt.Sprintf("%s\t%s", d.FileName, d.Name)
		}), cobra.ShellCompDirectiveNoFileComp
	}
}
