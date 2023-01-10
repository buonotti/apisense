package daemon

import (
	"golang.org/x/sys/unix"
)

// signals for the unix platform
const (
	SIGHUP  = unix.SIGHUP
	SIGINT  = unix.SIGINT
	SIGTERM = unix.SIGTERM
)
