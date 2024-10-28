package cmd

import (
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/validation/validators/pkg"
	"github.com/spf13/cobra"
)

var templatesUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update validator templates",
	Long:  "Update validator templates from the specified repos. Also discovers new repos from the official remote",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		err := pkg.Update(force)
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
		log.DefaultLogger().Info("Updated template repositories")
	},
}

func init() {
	templatesUpdateCmd.Flags().BoolP("force", "f", false, "Override the local repos with the online discovered ones from the official repo")

	templatesCmd.AddCommand(templatesUpdateCmd)
}
