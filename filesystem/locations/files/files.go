package files

import (
	"path/filepath"

	"github.com/buonotti/apisense/filesystem/locations/directories"
)

func DotenvFile() string {
	return filepath.FromSlash(directories.AppDirectory() + "/.env")
}

func DbFile() string {
	return filepath.FromSlash(directories.AppDirectory() + "/apisense.sqlite")
}

func PkgLockFile() string {
	return filepath.FromSlash(directories.ValidatorsDirectory() + "/repos.lock.json")
}
