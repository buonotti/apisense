package cmd

import (
	"github.com/auribuo/stylishcobra"
	"github.com/buonotti/apisense/config"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/charmbracelet/lipgloss"
	clog "github.com/charmbracelet/log"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "apisense",
	Short: "apisense is a tool to monitor data from a REST web service",
	Long: `This cli is used to start and interface with the apisense daemon. The daemon is used to monitor data from a REST web service.
There are multiple subcommands that can be used to interact with the daemon. For more information about a specific subcommand use the --help flag.`,
	Args:    cobra.NoArgs,
	Version: "2.0.0",
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
	Run: func(cmd *cobra.Command, _ []string) {
		cobra.CheckErr(cmd.Help())
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
	underline := lipgloss.NewStyle().Underline(true)
	bold := lipgloss.NewStyle().Bold(true)
	italic := lipgloss.NewStyle().Italic(true)
	ubold := underline.Inherit(bold)

	stylishcobra.Setup(rootCmd).
		DisableExtraNewlines().
		StyleHeadings(ubold.Foreground(lipgloss.ANSIColor(termenv.ANSIBlue))).
		StyleCommands(bold.Foreground(lipgloss.ANSIColor(termenv.ANSICyan))).
		StyleCmdShortDescr(lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(termenv.ANSIBrightYellow))).
		StyleAliases(italic).
		StyleExecName(bold.Foreground(lipgloss.ANSIColor(termenv.ANSICyan))).
		StyleExample(italic).
		StyleFlags(lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(termenv.ANSICyan))).
		StyleFlagsDescr(italic.Foreground(lipgloss.ANSIColor(termenv.ANSIBrightYellow))).
		StyleFlagsDataType(italic.Foreground(lipgloss.AdaptiveColor{
			Light: "#444444",
			Dark:  "#777777",
		})).
		Init()

	rootCmd.SetVersionTemplate("{{.Version}}\n")
}
