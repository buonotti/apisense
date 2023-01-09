package daemon

import (
	"golang.org/x/sys/unix"
)

const (
	sighup = unix.SIGHUP
	sigint = unix.SIGINT
)
