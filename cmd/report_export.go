package cmd

import (
	"fmt"
	"github.com/buonotti/apisense/log"

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
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}

		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}

		reports, err := pipeline.Reports()
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}

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
				if err != nil {
					log.DefaultLogger().Fatal(err)
				}
				return
			}

			converter := conversion.Get(format)
			if converter == nil {
				if err != nil {
					log.DefaultLogger().Fatal(err)
				}
			}

			str, err := converter.Convert(*report)
			if err != nil {
				log.DefaultLogger().Fatal(err)
			}
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
	if err != nil {
		log.DefaultLogger().Fatal(errors.CannotRegisterCompletionFunction.Wrap(err, "cannot register completion function for reports"))
	}
	err = reportExportCmd.MarkFlagRequired("format")
	if err != nil {
		log.DefaultLogger().Fatal(err)
	}
	reportCmd.AddCommand(reportExportCmd)
}
