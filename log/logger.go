package log

import (
	"io"
	"os"

	"github.com/charmbracelet/log"
)

// getWriter returns the io.Writer used to output information
func getWriter() io.Writer {
	if hasLogFile() {
		return logFile
	} else {
		return os.Stderr
	}
}

// DefaultLogger returns a logger with no prefix
func DefaultLogger() *log.Logger {
	return log.Default()
}

// DaemonLogger returns a logger with the prefix "Daemon"
func DaemonLogger() *log.Logger {
	return log.Default().WithPrefix("Daemon")
}

// ApiLogger returns a logger with the prefix "Api"
func ApiLogger() *log.Logger {
	return log.Default().WithPrefix("Api")
}

// SshLogger returns a logger with the prefix "Ssh"
func SshLogger() *log.Logger {
	return log.Default().WithPrefix("Ssh")
}

// TuiLogger returns a logger with the prefix "Tui"
func TuiLogger() *log.Logger {
	return log.Default().WithPrefix("Tui")
}

// PkgLogger returns a logger with the prefix "Pkg"
func PkgLogger() *log.Logger {
	return log.Default().WithPrefix("Pkg")
}
