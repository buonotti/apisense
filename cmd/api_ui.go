package cmd

import (
	"fmt"

	"github.com/buonotti/apisense/api"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/util"
	"github.com/spf13/cobra"
)

var apiUiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Manage the apisense web ui",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		install, err := cmd.Flags().GetBool("install")
		if err != nil {
			install = false
		}
		if install {
			errors.CheckErr(api.InstallUI())
		} else {
			if util.Exists(directories.UiDirectory()) {
				fmt.Printf("Ui is %s\n", greenStyle().Render("enabled"))
			} else {
				fmt.Printf("Ui is %s. Install it with: apisense api ui --install\n", redStyle().Render("disabled"))
			}
		}
	},
}

func init() {
	apiUiCmd.Flags().Bool("install", false, fmt.Sprintf("Install the ui to: %s", directories.UiDirectory()))

	apiCmd.AddCommand(apiUiCmd)
}
