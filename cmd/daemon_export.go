package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/conversion"
	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/errors"
)

var daemonExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export the currently available pipeline items",
	Long:  `This command exports the currently available pipeline items. The exported items are the ones that are available in the pipeline.`,
	Run: func(cmd *cobra.Command, args []string) {
		pipeline, err := daemon.CreatePipeline()
		errors.HandleError(err)
		format, err := cmd.Flags().GetString("format")
		errors.HandleError(errors.SafeWrap(errors.CannotGetFlagValueError, err, "Cannot get value of flag: format"))
		switch format {
		case "json":
			data, _ := json.Marshal(pipeline)
			fmt.Println(string(data))
		default:
			fmt.Println("Unknown format") // TODO csv?
		}
	},
}

func init() {
	daemonExportCmd.Flags().StringP("format", "f", "", "Specify the export format")
	err := daemonExportCmd.RegisterFlagCompletionFunc("format", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return conversion.Converters(), cobra.ShellCompDirectiveNoFileComp
	})
	errors.HandleError(errors.SafeWrap(errors.CannotRegisterCompletionFunction, err, "Cannot register completion function for daemon export"))
	errors.HandleError(daemonExportCmd.MarkFlagRequired("format"))
	daemonCmd.AddCommand(daemonExportCmd)
}
