package config

import (
	"os"
)

var GetAsset func(string) ([]byte, error)

func create() error {
	err := os.Mkdir(Directory, 0755)
	if err != nil {
		return err
	}
	data, err := GetAsset("assets/config.example.toml")
	if err != nil {
		return err
	}
	return os.WriteFile(FullPath, data, os.ModePerm)
}
