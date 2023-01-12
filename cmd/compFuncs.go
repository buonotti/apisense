package cmd

import (
	"github.com/spf13/cobra"

	"github.com/buonotti/odh-data-monitor/util"
	"github.com/buonotti/odh-data-monitor/validation"
)

func validReportsFunc() func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		reports, _ := validation.Reports()
		return util.Map(reports, func(r validation.Report) string { return r.Id }), cobra.ShellCompDirectiveNoFileComp
	}
}
