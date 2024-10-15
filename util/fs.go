package util

import (
	"os"
	"strings"
)

// ExpandHome replaes the character "~" in the string with the content of the "HOME" environment variable in place
func ExpandHome(path *string) {
	derefPath := *path
	if strings.Contains(derefPath, "~") {
		derefPath = strings.ReplaceAll(derefPath, "~", os.Getenv("HOME"))
	}
	*path = derefPath
}

// Exists returns true if the given file exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
