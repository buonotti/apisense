package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
)

var reportCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean the report directory",
	Long:  `This command cleans the report directory.`,
	Run: func(cmd *cobra.Command, _ []string) {
		override, err := cmd.Flags().GetBool("no-confirm")
		if err != nil {
			override = false
		}

		var answer string
		if !override {
			fmt.Print("Are you sure you want to clean the report directory? [y/N] ")
			_, err = fmt.Scanln(&answer)
			if err != nil {
				log.CliLogger.WithError(err).Fatal("cannot read user input")
				return
			}
		}
		if answer == "y" || answer == "Y" || override {
			log.CliLogger.Info("cleaning report directory")
			reportFiles, err := os.ReadDir(directories.ReportsDirectory())
			if err != nil {
				log.CliLogger.WithError(err).Fatal("cannot read report directory")
				return
			}

			for _, file := range reportFiles {
				err := os.Remove(filepath.FromSlash(directories.ReportsDirectory() + "/" + file.Name()))
				if err != nil {
					log.CliLogger.WithError(err).Fatal("cannot remove file")
					return
				}
				log.CliLogger.WithField("file", file.Name()).Info("removed file")
			}
		} else {
			log.CliLogger.Info("aborted")
		}
	},
}

func init() {
	reportCleanCmd.Flags().Bool("no-confirm", false, "Do not ask for confirmation")
	reportCmd.AddCommand(reportCleanCmd)

}
