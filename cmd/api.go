package cmd

import (
	"strconv"

	"github.com/buonotti/apisense/api"
	"github.com/buonotti/apisense/api/db"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start the api server",
	Long: `This command starts the api server. By default it runs on port 8080. and listens on all hosts.
The port and interface can be changed with the --port and --host flags. The flags override the values in the config file.`,
	Args: cobra.NoArgs,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmd.Root().PersistentPreRun(cmd, args)
		if viper.GetBool("api.auth") {
			cobra.CheckErr(db.Setup())
		}
	},
	Run: func(cmd *cobra.Command, _ []string) {
		host := cmd.Flag("host").Value.String()
		port := cmd.Flag("port").Value.String()
		portParsed, err := strconv.Atoi(port)
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
		err = api.Start(host, portParsed)
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
	},
}

func init() {
	apiCmd.Flags().String("host", "", "The host to listen on")
	apiCmd.Flags().Int("port", 8080, "The port to listen on")

	rootCmd.AddCommand(apiCmd)
}
