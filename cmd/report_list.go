package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/util"
	"github.com/buonotti/odh-data-monitor/validation"
)

var reportListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all available reports",
	Long:    "", // TODO
	Run: func(cmd *cobra.Command, args []string) {
		verbose, err := cmd.Flags().GetBool("verbose")
		errors.HandleError(err)
		reports, err := validation.Reports()
		errors.HandleError(err)
		reportIds := util.Map(reports, func(in validation.Report) string {
			if verbose {
				return fmt.Sprintf("%s at %s with %d result(s)", in.Id, in.Time.String(), len(in.Results))
			}
			return in.Id
		})
		for _, id := range reportIds {
			fmt.Println(id)
		}
	},
}

func init() {
	reportListCmd.Flags().BoolP("verbose", "v", false, "Be more verbose")
	reportCmd.AddCommand(reportListCmd)
}
