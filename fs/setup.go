package fs

import (
	"os"

	"github.com/buonotti/apisense/config"
	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/validation"
)

// Setup creates the daemon directory and writes the default files to it.
func Setup() error {
	err, definitionsExisted := createDirectories()
	if err != nil {
		return err
	}

	err = writeFiles(!definitionsExisted)
	if err != nil {
		return err
	}

	return nil
}

// createDirectories creates the daemon, reports and the definitions directories.
// If any of the directories exist it does nothing.
// If there is an error while creating the directories it returns an *errors.CannotCreateDirectoryError.
func createDirectories() (error, bool) {
	dataDirectory := os.Getenv("HOME") + "/apisense"
	definitionsExist := true

	if _, err := os.Stat(dataDirectory); os.IsNotExist(err) {
		err = os.Mkdir(dataDirectory, os.ModePerm)
		if err != nil {
			return errors.CannotCreateDirectoryError.Wrap(err, "cannot create apisense directory: "+dataDirectory), false
		}
	}

	if _, err := os.Stat(daemon.Directory()); os.IsNotExist(err) {
		if err = os.Mkdir(daemon.Directory(), 0755); err != nil {
			return errors.CannotCreateDirectoryError.Wrap(err, "cannot create daemon directory: "+daemon.Directory()), false
		}
	}

	if _, err := os.Stat(validation.DefinitionsLocation()); os.IsNotExist(err) {
		definitionsExist = false
		if err = os.Mkdir(validation.DefinitionsLocation(), 0755); err != nil {
			return errors.CannotCreateDirectoryError.Wrap(err, "cannot create definitions directory: "+validation.DefinitionsLocation()), false
		}
	}

	if _, err := os.Stat(validation.ReportLocation()); os.IsNotExist(err) {
		if err = os.Mkdir(validation.ReportLocation(), 0755); err != nil {
			return errors.CannotCreateDirectoryError.Wrap(err, "cannot create reports directory: "+validation.ReportLocation()), false
		}
	}

	return nil, definitionsExist
}

// writeFiles writes the default definition file to the definition directory.
// If there is an error while reading the default asset with config.GetAsset it return an *errors.CannotLoadAssetError.
// If there is an error while writing the file it returns an *errors.CannotWriteFileError.
func writeFiles(writeDefinitions bool) error {
	if _, err := os.Stat(daemon.StatusFile); os.IsNotExist(err) {
		err = os.WriteFile(daemon.StatusFile, []byte("down"), os.ModePerm)
		if err != nil {
			return errors.CannotWriteFileError.Wrap(err, "cannot write status file")
		}
	}

	if _, err := os.Stat(daemon.PidFile); os.IsNotExist(err) {
		err = os.WriteFile(daemon.PidFile, []byte("0"), os.ModePerm)
		if err != nil {
			return errors.CannotWriteFileError.Wrap(err, "cannot write pid file")
		}
	}

	if !writeDefinitions {
		return nil
	}

	defData, err := config.Asset("assets/bluetooth.toml")
	if err != nil {
		return errors.CannotLoadAssetError.Wrap(err, "cannot load bluetooth definition asset")
	}

	err = os.WriteFile(validation.DefinitionsLocation()+"/bluetooth.toml", defData, os.ModePerm)
	if err != nil {
		return errors.CannotWriteFileError.Wrap(err, "cannot write bluetooth definition file")
	}

	defData2, err := config.Asset("assets/bluetooth2.toml")
	if err != nil {
		return errors.CannotLoadAssetError.Wrap(err, "cannot load bluetooth2 definition asset")
	}

	err = os.WriteFile(validation.DefinitionsLocation()+"/bluetooth2.toml", defData2, os.ModePerm)
	if err != nil {
		return errors.CannotWriteFileError.Wrap(err, "cannot write bluetooth2 definition file")
	}

	return nil
}
