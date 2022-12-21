package config

import (
	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/spf13/viper"
)

var configName = "config"
var configDir = configdir.LocalConfig("odh-data-monitor")
var configPath = filepath.FromSlash(configDir + "/" + configName + ".toml")

func Setup() error {
	viper.SetConfigName(configName)
	viper.AddConfigPath(configDir)
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
