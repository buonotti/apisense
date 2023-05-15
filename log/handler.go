package log

import (
	"io"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/apex/log/handlers/json"
)

type dynamicHandler struct {
	defaultHandler log.Handler
	cliHandler     log.Handler
}

func newHandler() log.Handler {
	w := getWriter()

	cliHandler := cli.New(w)
	cliHandler.Padding = 0

	return &dynamicHandler{
		defaultHandler: json.New(w),
		cliHandler:     cliHandler,
	}
}

func getWriter() io.Writer {
	if hasLogFile() {
		return logFile
	} else {
		return os.Stdout
	}
}

func (h *dynamicHandler) HandleLog(e *log.Entry) error {
	if e.Fields["system"] == "cli" {
		return h.cliHandler.HandleLog(e)
	}

	if hasLogFile() {
		return h.defaultHandler.HandleLog(e)
	}
	return h.cliHandler.HandleLog(e)
}
