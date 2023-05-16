package cmd

import "github.com/spf13/cobra"

var apiUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
	Long:  `This command allows to manage the users of the API.`,
	Run: func(cmd *cobra.Command, _ []string) {
		cobra.CheckErr(cmd.Help())
	},
}

func init() {
	apiCmd.AddCommand(apiUserCmd)
}
