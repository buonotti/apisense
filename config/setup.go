package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/buonotti/apisense/v2/errors"
	"github.com/buonotti/apisense/v2/filesystem/locations/directories"
	"github.com/buonotti/apisense/v2/util"
	"github.com/spf13/viper"
)

const (
	FileName    string = "apisense.config"
	SecretsName string = "apisense.secrets"
)

// Setup loads the config file with viper.ReadInConfig and creates the default config if it doesn't exist.
func Setup() error {
	configFile := filepath.FromSlash(directories.ConfigDirectory() + "/" + FileName + ".yml")
	secretsFile := filepath.FromSlash(directories.ConfigDirectory() + "/" + SecretsName + ".yml")

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

		file.Close()
	}

	viper.SetEnvPrefix("apisense")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.AddConfigPath(filepath.FromSlash(directories.ConfigDirectory()))
	viper.SetConfigName(FileName)
	err = viper.ReadInConfig()
	if err != nil {
		return errors.CannotReadInConfigError.Wrap(err, "cannot read main config file")
	}

	if util.Exists(secretsFile) {
		SecretsViper.SetEnvPrefix("apisense_secret")
		SecretsViper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		SecretsViper.AutomaticEnv()

		SecretsViper.AddConfigPath(filepath.FromSlash(directories.ConfigDirectory()))
		SecretsViper.SetConfigName(SecretsName)
		err = SecretsViper.ReadInConfig()
		if err != nil {
			return errors.CannotReadInConfigError.Wrap(err, "cannot read secrets config file")
		}

		SecretsViper.WatchConfig()
	}

	viper.WatchConfig()

	return nil
}

var SecretsViper = viper.New()
