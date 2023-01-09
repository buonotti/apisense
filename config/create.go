package config

import (
	"os"

	"github.com/buonotti/odh-data-monitor/errors"
)

var GetAsset func(string) ([]byte, error)

func createExampleConfig() error {
	err := os.Mkdir(Directory, 0755)
	if err != nil {
		return errors.CannotCreateDirectoryError.Wrap(err, "Cannot create config directory")
	}
	data, err := GetAsset("assets/config.example.toml")
	if err != nil {
		return errors.CannotLoadAssetError.Wrap(err, "Cannot load config example asset")
	}
	err = os.WriteFile(FullPath, data, os.ModePerm)
	if err != nil {
		return errors.CannotWriteFileError.Wrap(err, "Cannot write config file")
	}
	return nil
}
