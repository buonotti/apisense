package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "One-shot validation",
	Long:  "Validate all the definitions only once, print the report then exit",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		validationPipeline, err := daemon.NewPipeline()
		if err != nil {
			log.DefaultLogger().Error("Failed to create pipeline", "reason", err.Error())
			os.Exit(1)
		}
		report := validationPipeline.Validate()
		data, err := json.Marshal(report)
		errors.CheckErr(err)
		fmt.Print(string(data))
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
