package log

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/apex/log"
	fColor "github.com/fatih/color"
	"github.com/mattn/go-colorable"
)

var bold = fColor.New(fColor.Bold)

// Colors mapping.
var Colors = [...]*fColor.Color{
	log.DebugLevel: fColor.New(fColor.FgWhite),
	log.InfoLevel:  fColor.New(fColor.FgBlue),
	log.WarnLevel:  fColor.New(fColor.FgYellow),
	log.ErrorLevel: fColor.New(fColor.FgRed),
	log.FatalLevel: fColor.New(fColor.FgRed),
}

// Strings mapping.
var Strings = [...]string{
	log.DebugLevel: "•",
	log.InfoLevel:  "•",
	log.WarnLevel:  "!",
	log.ErrorLevel: "⨯",
	log.FatalLevel: "⨯",
}

// cliHandler implementation.
type cliHandler struct {
	mu      sync.Mutex
	Writer  io.Writer
	Padding int
}

// newCliHandler handler.
func newCliHandler(w io.Writer) *cliHandler {
	if f, ok := w.(*os.File); ok {
		return &cliHandler{
			Writer:  colorable.NewColorable(f),
			Padding: 0,
		}
	}

	return &cliHandler{
		Writer:  w,
		Padding: 0,
	}
}

// HandleLog implements log.cliHandler.
func (h *cliHandler) HandleLog(e *log.Entry) error {
	color := Colors[e.Level]
	level := Strings[e.Level]

	h.mu.Lock()
	defer h.mu.Unlock()

	_, err := color.Fprintf(h.Writer, "%s %-25s", bold.Sprintf("%*s", h.Padding+1, level), e.Message)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(h.Writer)
	if err != nil {
		return err
	}

	return nil
}
