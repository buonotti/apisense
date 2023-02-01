package util

import (
	"strings"
)

func Pad(s string, maxLen int) string {
	if len(s) < maxLen {
		return s + strings.Repeat(" ", maxLen-len(s))
	}
	return s
}
