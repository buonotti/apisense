package daemon

import (
	"os"
)

// Directory is the path of the directory containing the daemon control files
func Directory() string {
	return os.Getenv("HOME") + "/odh-data-monitor/daemon"
}
