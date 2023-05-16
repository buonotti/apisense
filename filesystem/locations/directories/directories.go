package directories

import (
	"path/filepath"

	"github.com/auribuo/userdir"
)

var app userdir.Userdir = userdir.App("apisense")

func AppDirectory() string {
	dataDir, _ := app.Data()
	return filepath.ToSlash(dataDir)
}

func ConfigDirectory() string {
	configDir, err := app.Config()
	if err != nil {
		return AppDirectory()
	}
	return filepath.ToSlash(configDir)
}
