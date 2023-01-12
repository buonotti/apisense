package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/buonotti/odh-data-monitor/conversion"
	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/util"
	"github.com/buonotti/odh-data-monitor/validation"
)

var reportExportCmd = &cobra.Command{
	Use:   "export [flags] [reports]...",
	Short: "Export reports in various formats",
	Long:  "", // TODO
	Run: func(cmd *cobra.Command, args []string) {
		format, err := cmd.Flags().GetString("format")
		errors.HandleError(errors.SafeWrap(errors.CannotGetFlagValueError, err, "Cannot get value for flag: format"))

		all, err := cmd.Flags().GetBool("all")
		errors.HandleError(errors.SafeWrap(errors.CannotGetFlagValueError, err, "Cannot get value for flag: all"))

		reports, err := validation.Reports()
		errors.HandleError(err)

		var ids []string
		if all {
			ids = util.Map(reports, func(r validation.Report) string { return r.Id })
		} else {
			ids = args
		}

		for _, arg := range ids {
			report := util.FindFirst(reports, func(report validation.Report) bool {
				return report.Id == arg
			})
			if report != nil {
				errors.HandleError(errors.NewF(errors.UnknownReportError, "Unknown report: %s", arg))
			}
			switch format {
			case "json":
				str, err := conversion.Json().Convert(*report)
				errors.HandleError(err)
				fmt.Println(string(str))
			case "csv":
				str, err := conversion.Csv().Convert(*report)
				errors.HandleError(err)
				fmt.Println(string(str))
			default:
				errors.HandleError(errors.NewF(errors.UnknownExportFormatError, "Unknown format: %s", format))
			}
		}
	},
	Args:              cobra.MinimumNArgs(1),
	ValidArgsFunction: validReportsFunc(),
}

func init() {
	reportExportCmd.Flags().StringP("format", "f", "", "Specify the export format")
	reportExportCmd.Flags().Bool("all", false, "Export all reports")
	err := reportExportCmd.RegisterFlagCompletionFunc("format", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"json", "csv"}, cobra.ShellCompDirectiveNoFileComp
	})
	errors.HandleError(errors.SafeWrap(errors.CannotRegisterCompletionFunction, err, "Cannot register completion function for reports"))
	errors.HandleError(reportExportCmd.MarkFlagRequired("format"))
	reportCmd.AddCommand(reportExportCmd)
}
