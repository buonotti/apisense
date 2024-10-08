package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/errors"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "One-shot validation",
	Long:  "Validate all the definitions only once, print the report then exit",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		validationPipeline, err := daemon.NewPipeline()
		errors.CheckErr(err)
		report, err := validationPipeline.Validate()
		errors.CheckErr(err)
		data, err := json.Marshal(report)
		errors.CheckErr(err)
		fmt.Println(string(data))
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
