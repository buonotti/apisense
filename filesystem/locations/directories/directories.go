package directories

import (
	"os"
	"path/filepath"
)

func AppDirectory() string {
	return filepath.ToSlash(os.Getenv("HOME")) + "/apisense"
}

func ConfigDirectory() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return AppDirectory()
	}
	return filepath.ToSlash(configDir) + "/apisense"
}
