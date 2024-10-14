package log

import (
	"io"
	"os"

	"github.com/charmbracelet/log"
)

func getWriter() io.Writer {
	if hasLogFile() {
		return logFile
	} else {
		return os.Stderr
	}
}

func DefaultLogger() *log.Logger {
	return log.Default()
}

func DaemonLogger() *log.Logger {
	return log.Default().WithPrefix("Daemon")
}

func ApiLogger() *log.Logger {
	return log.Default().WithPrefix("Api")
}

func SshLogger() *log.Logger {
	return log.Default().WithPrefix("Ssh")
}

func TuiLogger() *log.Logger { return log.Default().WithPrefix("Tui") }
