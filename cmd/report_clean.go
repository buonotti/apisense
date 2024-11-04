package cmd

import (
	"os"
	"path/filepath"

	"github.com/buonotti/apisense/v2/errors"
	"github.com/buonotti/apisense/v2/filesystem/locations/directories"
	"github.com/buonotti/apisense/v2/log"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var reportCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean the report directory",
	Long:  `This command cleans the report directory.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		override, err := cmd.Flags().GetBool("no-confirm")
		if err != nil {
			override = false
		}

		var answer bool = true
		if !override {
			err = huh.
				NewConfirm().
				Title("Are you sure you want to clean the report directory?").
				Affirmative("Yes").
				Negative("No").
				Value(&answer).
				WithTheme(huh.ThemeCatppuccin()).
				Run()
			if err != nil {
				log.DefaultLogger().Fatal(errors.CannotRunPromptError.WrapWithNoMessage(err))
			}
		}
		if answer {
			log.DefaultLogger().Info("cleaning report directory")
			reportFiles, err := os.ReadDir(directories.ReportsDirectory())
			if err != nil {
				log.DefaultLogger().Fatal(errors.CannotReadDirectoryError.WrapWithNoMessage(err))
			}

			for _, file := range reportFiles {
				err := os.Remove(filepath.FromSlash(directories.ReportsDirectory() + "/" + file.Name()))
				if err != nil {
					log.DefaultLogger().Fatal(errors.CannotRemoveFileError.WrapWithNoMessage(err))
				}
				log.DefaultLogger().Info("Removed file", "file", file.Name())
			}
		} else {
			log.DefaultLogger().Info("Aborted")
		}
	},
}

func init() {
	reportCleanCmd.Flags().Bool("no-confirm", false, "Do not ask for confirmation")
	reportCmd.AddCommand(reportCleanCmd)
}
