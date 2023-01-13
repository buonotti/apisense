package fs

import (
	"os"
	"path/filepath"

	"github.com/buonotti/odh-data-monitor/config"
	"github.com/buonotti/odh-data-monitor/daemon"
	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/validation"
)

// Setup creates the daemon directory and writes the default files to it.
func Setup() error {
	err := createDirectories()
	if err != nil {
		return err
	}

	err = writeFiles()
	if err != nil {
		return err
	}

	return nil
}

// createDirectories creates the daemon, reports and the definitions directories.
// If any of the directories exist it does nothing.
// If there is an error while creating the directories it returns an *errors.CannotCreateDirectoryError.
func createDirectories() error {
	if _, err := os.Stat(os.Getenv("HOME") + "/odh-data-monitor"); os.IsNotExist(err) {
		err = os.Mkdir(os.Getenv("HOME")+"/odh-data-monitor", os.ModePerm)
		if err != nil {
			return errors.CannotCreateDirectoryError.Wrap(err, "Cannot create odh-data-monitor directory")
		}
	}

	if _, err := os.Stat(daemon.Directory()); os.IsNotExist(err) {
		if err = os.Mkdir(daemon.Directory(), 0755); err != nil {
			return errors.CannotCreateDirectoryError.Wrap(err, "Cannot create daemon directory")
		}
	}

	if _, err := os.Stat(validation.DefinitionsLocation()); os.IsNotExist(err) {
		if err = os.Mkdir(validation.DefinitionsLocation(), 0755); err != nil {
			return errors.CannotCreateDirectoryError.Wrap(err, "Cannot create definitions directory")
		}
	}

	if _, err := os.Stat(validation.ReportLocation()); os.IsNotExist(err) {
		if err = os.Mkdir(validation.ReportLocation(), 0755); err != nil {
			return errors.CannotCreateDirectoryError.Wrap(err, "Cannot create reports directory")
		}
	}

	return nil
}

// writeFiles writes the default definition file to the definition directory.
// If there is an error while reading the default asset with config.GetAsset it return an *errors.CannotLoadAssetError.
// If there is an error while writing the file it returns an *errors.CannotWriteFileError.
func writeFiles() error {
	defData, err := config.GetAsset("assets/bluetooth.definition.toml")
	if err != nil {
		return errors.CannotLoadAssetError.Wrap(err, "Cannot load bluetooth definition asset")
	}

	err = os.WriteFile(validation.DefinitionsLocation()+string(filepath.Separator)+"bluetooth.toml", defData, os.ModePerm)
	if err != nil {
		return errors.CannotWriteFileError.Wrap(err, "Cannot write bluetooth definition file")
	}

	if _, err := os.Stat(daemon.StatusFile); os.IsNotExist(err) {
		err = os.WriteFile(daemon.StatusFile, []byte("down"), os.ModePerm)
		if err != nil {
			return errors.CannotWriteFileError.Wrap(err, "Cannot write status file")
		}
	}

	if _, err := os.Stat(daemon.PidFile); os.IsNotExist(err) {
		err = os.WriteFile(daemon.PidFile, []byte("0"), os.ModePerm)
		if err != nil {
			return errors.CannotWriteFileError.Wrap(err, "Cannot write pid file")
		}
	}

	return nil
}
