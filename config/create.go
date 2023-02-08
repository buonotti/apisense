package config

import (
	"os"

	"github.com/buonotti/apisense/errors"
)

// create creates the config directory in the user config directory and writes the default config file to it.
func create() error {
	err := os.Mkdir(Directory, 0755)
	if err != nil && !os.IsExist(err) {
		return errors.CannotCreateDirectoryError.Wrap(err, "cannot create config directory")
	}

	data, err := Asset("assets/config.example.toml")
	if err != nil {
		return errors.CannotLoadAssetError.Wrap(err, "cannot load config example asset")
	}

	err = os.WriteFile(FullPath, data, os.ModePerm)
	if err != nil {
		return errors.CannotWriteFileError.Wrap(err, "cannot write config file")
	}

	return nil
}

func createEnv() error {
	data, err := Asset("assets/.env.example")
	if err != nil {
		return errors.CannotLoadAssetError.Wrap(err, "cannot load .env example asset")
	}

	err = os.WriteFile(os.Getenv("HOME")+"/apisense/.env", data, os.ModePerm)
	if err != nil {
		return errors.CannotWriteFileError.Wrap(err, "cannot write .env file")
	}

	return nil
}
