package tui

import (
	"github.com/buonotti/apisense/errors"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func getTerminalHeight() int {
	_, terminalHeight, err := terminal.GetSize(int(os.Stdin.Fd()))
	errors.CheckErr(err)
	return terminalHeight
}
