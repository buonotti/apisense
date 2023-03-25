package tui

import (
	"os"

	"github.com/buonotti/apisense/errors"
	"golang.org/x/term"
)

func getTerminalHeight() int {
	_, terminalHeight, err := term.GetSize(int(os.Stdin.Fd()))
	errors.CheckErr(err)
	return terminalHeight
}
