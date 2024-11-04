package files

import (
	"path/filepath"

	"github.com/buonotti/apisense/v2/filesystem/locations/directories"
)

func DaemonStatusFile() string {
	return filepath.FromSlash(directories.DaemonDirectory() + "/status")
}

func DaemonPidFile() string {
	return filepath.FromSlash(directories.DaemonDirectory() + "/daemon.pid")
}

func DaemonLockFile() string {
	return filepath.FromSlash(directories.DaemonDirectory() + "/daemon.lock")
}
