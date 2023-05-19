package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/conversion"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation/pipeline"
)

var reportExportCmd = &cobra.Command{
	Use:   "export [FLAGS] [REPORTS]...",
	Short: "Export reports in various formats",
	Long:  "This command exports all the reports in the report directory in one of the specified formats.",
	Run: func(cmd *cobra.Command, args []string) {
		format, err := cmd.Flags().GetString("format")
		errors.CheckErr(errors.SafeWrap(errors.CannotGetFlagValueError, err, "cannot get value for flag: format"))

		all, err := cmd.Flags().GetBool("all")
		errors.CheckErr(errors.SafeWrap(errors.CannotGetFlagValueError, err, "cannot get value for flag: all"))

		reports, err := pipeline.Reports()
		errors.CheckErr(err)

		var ids []string
		if all {
			ids = util.Map(reports, func(r pipeline.Report) string { return r.Id })
		} else {
			ids = args
		}

		for _, arg := range ids {
			report := util.FindFirst(reports, func(report pipeline.Report) bool {
				return report.Id == arg
			})

			if report == nil {
				errors.CheckErr(errors.NewF(errors.UnknownReportError, "unknown report: %s", arg))
				return
			}

			converter := conversion.Get(format)
			if converter == nil {
				errors.CheckErr(errors.NewF(errors.UnknownExportFormatError, "unknown format: %s", format))
			}

			str, err := converter.Convert(*report)
			errors.CheckErr(err)
			fmt.Println(string(str))
		}
	},
	ValidArgsFunction: validReportsFunc(),
}

func init() {
	reportExportCmd.Flags().StringP("format", "f", "", "Specify the export format")
	reportExportCmd.Flags().Bool("all", false, "Export all reports")
	err := reportExportCmd.RegisterFlagCompletionFunc("format", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return conversion.Converters(), cobra.ShellCompDirectiveNoFileComp
	})
	errors.CheckErr(errors.SafeWrap(errors.CannotRegisterCompletionFunction, err, "cannot register completion function for reports"))
	errors.CheckErr(reportExportCmd.MarkFlagRequired("format"))
	reportCmd.AddCommand(reportExportCmd)
}
