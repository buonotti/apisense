package daemon

import (
	"golang.org/x/sys/unix"
)

const (
	SIGHUP  = unix.SIGHUP
	SIGINT  = unix.SIGINT
	SIGTERM = unix.SIGTERM
)
