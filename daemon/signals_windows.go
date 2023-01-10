package daemon

import (
	"golang.org/x/sys/windows"
)

// signals for the windows platform
const (
	SIGHUP  = windows.SIGHUP
	SIGINT  = windows.SIGINT
	SIGTERM = windows.SIGTERM
)
