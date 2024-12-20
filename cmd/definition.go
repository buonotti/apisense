package cmd

import (
	"github.com/spf13/cobra"
)

var definitionCmd = &cobra.Command{
	Use:   "definition",
	Short: "Manage definitions",
	Long:  `This command is used to manage the definition files.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		cobra.CheckErr(cmd.Help())
	},
}

func init() {
	rootCmd.AddCommand(definitionCmd)
}
