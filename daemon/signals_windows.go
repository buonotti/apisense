package daemon

import (
	"golang.org/x/sys/windows"
)

const (
	sighup = windows.SIGHUP
	sigint = windows.SIGINT
)
