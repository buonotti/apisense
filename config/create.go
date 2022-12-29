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
	err = os.Mkdir(Directory+"/definitions.d", 0755)
	if err != nil {
		return err
	}
	data, err := GetAsset("assets/config.example.toml")
	if err != nil {
		return err
	}
	err = os.WriteFile(FullPath, data, os.ModePerm)
	if err != nil {
		return err
	}
	data, err = GetAsset("assets/apiname.definition.toml")
	if err != nil {
		return err
	}
	err = os.WriteFile(Directory+"/definitions.d/apiname.definition.toml", data, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
