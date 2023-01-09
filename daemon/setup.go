package daemon

import (
	"os"
	"path/filepath"

	"github.com/buonotti/odh-data-monitor/config"
	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/validation"
)

func Directory() string {
	return config.Directory + string(filepath.Separator) + "daemon"
}

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

func createDirectories() error {
	if _, err := os.Stat(Directory()); os.IsNotExist(err) {
		err := os.Mkdir(Directory(), 0755)
		if err != nil {
			return errors.CannotCreateDirectoryError.Wrap(err, "Cannot create daemon directory")
		}
	}
	if _, err := os.Stat(validation.DefinitionsLocation()); os.IsNotExist(err) {
		err := os.Mkdir(validation.DefinitionsLocation(), 0755)
		if err != nil {
			return errors.CannotCreateDirectoryError.Wrap(err, "Cannot create definitions directory")
		}
	}
	return nil
}

func writeFiles() error {
	defData, err := config.GetAsset("assets/bluetooth.definition.toml")
	if err != nil {
		return errors.CannotLoadAssetError.Wrap(err, "Cannot load bluetooth definition asset")
	}
	err = os.WriteFile(validation.DefinitionsLocation()+string(filepath.Separator)+"bluetooth.toml", defData, os.ModePerm)
	if err != nil {
		return errors.CannotWriteFileError.Wrap(err, "Cannot write bluetooth definition file")
	}
	return nil
}
