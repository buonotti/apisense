package cmd

import (
	"fmt"
	"time"

	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation/pipeline"
	"github.com/spf13/cobra"
)

var reportListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all available reports",
	Long:    "This command lists all available reports.",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
		reports, err := pipeline.Reports()
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
		reportIds := util.Map(reports, func(in pipeline.Report) string {
			if verbose {
				return fmt.Sprintf("%s --- %s with %d result(s)", in.Id, time.Time(in.Time).Format("2006-01-02 at 15-04-05.000Z"), len(in.Endpoints))
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
