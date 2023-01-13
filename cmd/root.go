package cmd

import (
	"github.com/spf13/cobra"

	cc "github.com/ivanpirog/coloredcobra"

	"github.com/buonotti/odh-data-monitor/config"
	"github.com/buonotti/odh-data-monitor/fs"
	"github.com/buonotti/odh-data-monitor/log"
)

var rootCmd = &cobra.Command{
	Use:   "odh-data-monitor",
	Short: "odh-data-monitor is a tool to monitor data from the Open Data Hub",
	Long: `This program provides multiple ways to interface with the daemon that is monitoring the Open Data Hub apis.
In the first place this CLI can be used to start the daemon itself. It also provides a TUI to manage the daemon and its configs
in a more user friendly way. The program can also start an SSH-Server that serves the tui over SSH and automatically starts the
daemon. For more info check each commands description.`,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(cmd.Help())
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(fs.Setup())
		cobra.CheckErr(config.Setup())
		cobra.CheckErr(log.Setup())
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
