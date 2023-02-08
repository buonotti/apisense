package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/errors"
)

var daemonExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export the currently available pipeline items",
	Long:  `This command exports the currently available pipeline items. The exported items are the ones that are available in the pipeline.`,
	Run: func(cmd *cobra.Command, args []string) {
		pipeline, err := daemon.NewPipeline()
		errors.CheckErr(err)
		format, err := cmd.Flags().GetString("format")
		errors.CheckErr(errors.SafeWrap(errors.CannotGetFlagValueError, err, "cannot get value of flag: format"))
		switch format {
		case "json":
			data, _ := json.Marshal(pipeline)
			fmt.Println(string(data))
		default:
			errors.CheckErr(errors.UnknownFormatError.New("invalid format: %s", format))
		}
	},
}

func init() {
	daemonExportCmd.Flags().StringP("format", "f", "", "Specify the export format")
	err := daemonExportCmd.RegisterFlagCompletionFunc("format", validFormatsFunc())
	errors.CheckErr(errors.SafeWrap(errors.CannotRegisterCompletionFunction, err, "cannot register completion function for daemon export"))
	errors.CheckErr(daemonExportCmd.MarkFlagRequired("format"))
	daemonCmd.AddCommand(daemonExportCmd)
}
