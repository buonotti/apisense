package cmd

import (
	"github.com/spf13/cobra"

	cc "github.com/ivanpirog/coloredcobra"

	"github.com/buonotti/odh-data-monitor/config"
	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/log"
)

var rootCmd = &cobra.Command{
	Use:     "odh-data-monitor",
	Short:   "odm is a tool for managing the OpenDataHub Data Monitor",
	Long:    `odm is a tool for managing the OpenDataHub Data Monitor`, // TODO add more info
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(cmd.Help())
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		errors.HandleError(config.Setup())
		errors.HandleError(log.Setup())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
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
