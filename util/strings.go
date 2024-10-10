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

func Capitalize(s string) string {
	bldr := strings.Builder{}
	upper := strings.ToUpper(s[:1])
	rest := s[1:]
	bldr.WriteString(upper)
	bldr.WriteString(rest)
	return bldr.String()
}
