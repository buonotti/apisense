package config

import (
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/errors"
)

// FileName is the name of the config file without extension
var FileName = "config"

// Directory is the path of the users config directory
var Directory = configdir.LocalConfig("apisense")

// FullPath is the full path of the config file with extension
var FullPath = Directory + string(filepath.Separator) + FileName + ".toml"

// Setup loads the config file with viper.ReadInConfig and creates the default config if it doesn't exist.
func Setup() error {
	viper.SetConfigFile(os.Getenv("HOME") + "/apisense/.env")
	err := viper.ReadInConfig()
	if err != nil {
		err = createEnv()
		if err != nil {
			return err
		}
		err = viper.ReadInConfig()
		if err != nil {
			return errors.CannotReadInConfigError.Wrap(err, "cannot read in .env file")
		}
	}

	viper.SetConfigName(FileName)
	viper.AddConfigPath(Directory)
	err = viper.MergeInConfig()
	if err != nil {
		err = create()
		if err != nil {
			return err
		}

		err = viper.MergeInConfig()
		if err != nil {
			return errors.CannotReadInConfigError.Wrap(err, "cannot read in main config file")
		}
	}

	return nil
}
