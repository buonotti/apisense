package cmd

import (
	"os"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/log"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var nukeCmd = &cobra.Command{
	Use:   "nuke",
	Short: "Kaboom",
	Long:  "Delete everything apisense related",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var delete bool
		err := huh.
			NewConfirm().
			Title("You will lose all data related to apisense. All non saved data will be gone for good").
			Affirmative("NUKE!").
			Negative("ABORT MISSION!!!").
			Value(&delete).WithTheme(huh.ThemeCatppuccin()).
			Run()
		if err != nil {
			log.DefaultLogger().Fatal(err)
		}
		if delete {
			err = os.RemoveAll(directories.AppDirectory())
			if err != nil {
				log.DefaultLogger().Fatal(errors.CannotRemoveFileError.WrapWithNoMessage(err))
			}
			err = os.RemoveAll(directories.ConfigDirectory())
			if err != nil {
				log.DefaultLogger().Fatal(errors.CannotRemoveFileError.WrapWithNoMessage(err))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(nukeCmd)
}
