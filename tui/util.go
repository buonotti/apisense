package tui

import (
	tSize "github.com/kopoli/go-terminal-size"
)

func getTerminalHeight() int {
	s, _ := tSize.GetSize()
	return s.Height
}
