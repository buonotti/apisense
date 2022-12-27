package tui

import (
	"os"

	"golang.org/x/term"
)

func getTerminalWidth() int {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	return width
}

func getTerminalHeight() int {
	_, height, _ := term.GetSize(int(os.Stdout.Fd()))
	return height
}
