package cmd

import (
	"github.com/spf13/cobra"

	cc "github.com/ivanpirog/coloredcobra"

	"github.com/buonotti/apisense/config"
	"github.com/buonotti/apisense/filesystem"
	"github.com/buonotti/apisense/log"
)

var rootCmd = &cobra.Command{
	Use:   "apisense",
	Short: "apisense is a tool to monitor data from a REST web service",
	Long: `This cli is used to start and interface with the apisense daemon. The daemon is used to monitor data from a REST web service.
There are multiple subcommands that can be used to interact with the daemon. For more information about a specific subcommand use the --help flag.`,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, _ []string) {
		cobra.CheckErr(cmd.Help())
	},
	PersistentPreRun: func(_ *cobra.Command, _ []string) {
		cobra.CheckErr(filesystem.Setup())
		cobra.CheckErr(config.Setup())
		cobra.CheckErr(log.Setup())
	},
	PersistentPostRun: func(_ *cobra.Command, _ []string) {
		cobra.CheckErr(log.CloseLogFile())
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
