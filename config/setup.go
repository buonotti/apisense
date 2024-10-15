package config

import (
	"os"
	"path/filepath"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/filesystem/locations/files"
	"github.com/buonotti/apisense/util"
	"github.com/spf13/viper"
)

const (
	FileName    string = "apisense.config"
	SecretsName string = "apisense.secrets"
)

// Setup loads the config file with viper.ReadInConfig and creates the default config if it doesn't exist.
func Setup() error {
	configFile := filepath.FromSlash(directories.ConfigDirectory() + "/" + FileName + ".yml")

	err := os.MkdirAll(filepath.FromSlash(directories.ConfigDirectory()), 0o755)
	if err != nil {
		return errors.CannotCreateDirectoryError.Wrap(err, "cannot create config directory")
	}

	setupDefaults()

	if !util.Exists(configFile) {
		file, err := os.Create(configFile)
		if err != nil {
			return errors.CannotCreateFileError.Wrap(err, "cannot create config file")
		}

		defer file.Close()
	}

	viper.AddConfigPath(filepath.FromSlash(directories.ConfigDirectory()))
	viper.SetConfigName(FileName)
	err = viper.ReadInConfig()
	if err != nil {
		return errors.CannotReadInConfigError.Wrap(err, "cannot read main config file")
	}

	viper.SetConfigName(SecretsName)
	_ = viper.MergeInConfig()

	viper.WatchConfig()

	return nil
}

func setupEnv() error {
	viper.SetConfigFile(files.DotenvFile())
	err := viper.ReadInConfig()

	if err != nil && !os.IsNotExist(err) {
		return errors.CannotReadFileError.Wrap(err, "cannot read config file")
	}

	return nil
}
