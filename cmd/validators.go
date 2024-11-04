package cmd

import "github.com/spf13/cobra"

var validatorsCmd = &cobra.Command{
	Use:   "validators",
	Short: "Manage external validators",
	Long:  "This command is used to manage external validators installed to the default location", // TODO: apisense env
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(cmd.Help())
	},
}

func init() {
	rootCmd.AddCommand(validatorsCmd)
}
