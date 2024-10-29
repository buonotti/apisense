package cmd

import (
	"fmt"

	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/log"
	"github.com/goccy/go-json"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "One-shot validation",
	Long:  "Validate all the definitions only once, print the report then exit.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		validationPipeline, err := daemon.NewPipeline()
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
		report, err := validationPipeline.Validate()
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
		data, err := json.Marshal(report)
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
		fmt.Println(string(data))
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
