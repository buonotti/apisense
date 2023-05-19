package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/buonotti/apisense/api"
	"github.com/buonotti/apisense/errors"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start the api server",
	Long: `This command starts the api server. By default it runs on port 8080. and listens on all hosts.
The port and interface can be changed with the --port and --host flags. The flags override the values in the config file`,
	Run: func(cmd *cobra.Command, _ []string) {
		host := cmd.Flag("host").Value.String()
		port := cmd.Flag("port").Value.String()
		portParsed, err := strconv.Atoi(port)
		errors.CheckErr(err)
		errors.CheckErr(api.Start(host, portParsed))
	},
}

func init() {
	apiCmd.Flags().String("host", "", "The host to listen on")
	apiCmd.Flags().Int("port", 8080, "The port to listen on")
	rootCmd.AddCommand(apiCmd)
}
