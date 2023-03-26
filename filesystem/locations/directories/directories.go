package directories

import (
	"os"
	"path/filepath"
)

func AppDirectory() string {
	home, _ := os.UserHomeDir()
	return filepath.ToSlash(home) + "/apisense"
}

func ConfigDirectory() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return AppDirectory()
	}
	return filepath.ToSlash(configDir) + "/apisense"
}
