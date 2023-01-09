package config

import (
	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/spf13/viper"
)

var FileName = "config"
var Directory = configdir.LocalConfig("odh-data-monitor")
var FullPath = Directory + string(filepath.Separator) + FileName + ".toml"

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
