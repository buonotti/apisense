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

func CacheDirectory() string {
	cacheDir, _ := app.Cache()
	return filepath.ToSlash(cacheDir)
}

func ValidatorsDirectory() string {
	return AppDirectory() + "/validators"
}

func ValidatorRepoDirectory() string {
	return ValidatorsDirectory() + "/repo"
}

func ValidatorCustomDirectory() string {
	return ValidatorsDirectory() + "/custom"
}

func ValidatorsCacheDirectory() string {
	return CacheDirectory() + "/validators"
}
