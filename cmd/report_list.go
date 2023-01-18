package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation"
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
				return fmt.Sprintf("%s --- %s with %d result(s)", in.Id, time.Time(in.Time).Format("2006-01-02 at 15-04-05.000Z"), len(in.Results))
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
