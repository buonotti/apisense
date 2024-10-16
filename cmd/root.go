package cmd

import (
	"github.com/buonotti/apisense/config"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	clog "github.com/charmbracelet/log"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "apisense",
	Short: "apisense is a tool to monitor data from a REST web service",
	Long: `This cli is used to start and interface with the apisense daemon. The daemon is used to monitor data from a REST web service.
There are multiple subcommands that can be used to interact with the daemon. For more information about a specific subcommand use the --help flag.`,
	Args:    cobra.NoArgs,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, _ []string) {
		cobra.CheckErr(cmd.Help())
	},
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		cobra.CheckErr(config.Setup())
		cobra.CheckErr(log.Setup())

		if ll, err := cmd.Root().PersistentFlags().GetString("log-level"); err == nil && ll != "" {
			parsed, err := clog.ParseLevel(ll)
			if err == nil {
				clog.SetLevel(parsed)
			} else {
				log.DefaultLogger().Warn("Log level invalid. Falling back to config", "reason", err.Error())
			}
		}
	},
	PersistentPostRun: func(_ *cobra.Command, _ []string) {
		cobra.CheckErr(log.CloseLogFile())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.PersistentFlags().String("log-level", "", "Override the log level in the config")
	err := rootCmd.RegisterFlagCompletionFunc("log-level", validLogLevelsFunc())
	if err != nil {
		log.DefaultLogger().Fatal(errors.CannotRegisterCompletionFunction.WrapWithNoMessage(err))
	}
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cc.Init(&cc.Config{
		RootCmd:       rootCmd,
		Headings:      cc.HiCyan + cc.Bold + cc.Underline,
		Commands:      cc.HiYellow + cc.Bold,
		Example:       cc.Italic,
		ExecName:      cc.Bold,
		Flags:         cc.Bold,
		FlagsDataType: cc.Italic + cc.HiBlue,
	})
	rootCmd.SetVersionTemplate("{{.Version}}\n")
}
