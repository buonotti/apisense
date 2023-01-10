package config

import (
	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/spf13/viper"
)

// FileName is the name of the config file without extension
var FileName = "config"

// Directory is the path of the users config directory
var Directory = configdir.LocalConfig("odh-data-monitor")

// FullPath is the full path of the config file with extension
var FullPath = Directory + string(filepath.Separator) + FileName + ".toml"

// Setup loads the config file with viper.ReadInConfig and creates the default config if it doesn't exist.
func Setup() error {
	viper.SetConfigName(FileName)
	viper.AddConfigPath(Directory)
	err := viper.ReadInConfig()
	if err != nil {
		err = create()
		if err != nil {
			return err
		}
		err = viper.ReadInConfig()
		if err != nil {
			return err
		}
	}
	return nil
}
