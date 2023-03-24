package config

import (
	"os"
	"path/filepath"

	"github.com/buonotti/apisense/errors"
	"github.com/kirsle/configdir"
	"github.com/spf13/viper"
)

// FileName is the name of the config file without extension
var FileName = "apisense.config"

// Location is the path of the users config directory
var Location = filepath.ToSlash(configdir.LocalConfig("apisense"))

// FullPath is the full path of the config file with extension
// var FullPath = Location + string(filepath.Separator) + FileName + ".yml"

// Setup loads the config file with viper.ReadInConfig and creates the default config if it doesn't exist.
func Setup() error {
	if _, err := os.Stat(Location); os.IsNotExist(err) {
		err = os.MkdirAll(Location, 0755)
		if err != nil {
			return errors.CannotCreateDirectoryError.Wrap(err, "cannot create config directory")
		}
	}

	setupDefaults()

	err := setupEnv()

	if err != nil {
		return err
	}

	viper.SetConfigName(FileName)
	viper.AddConfigPath(Location)

	err = viper.MergeInConfig()

	if err != nil {
		// err = create()
		// if err != nil {
		//	return err
		// }

		// err = viper.MergeInConfig()
		// if err != nil {
		//	 return errors.CannotReadInConfigError.Wrap(err, "cannot read in main config file")
		// }
	}

	return nil
}

func setupEnv() error {
	viper.SetConfigFile(os.Getenv("HOME") + "/apisense/.env")
	err := viper.ReadInConfig()

	if err != nil {
		// err = createEnv()
		// if err != nil {
		//	return err
		// }
		// err = viper.ReadInConfig()
		// if err != nil {
		//	return errors.CannotReadInConfigError.Wrap(err, "cannot read in .env file")
		// }
	}

	return nil
}
