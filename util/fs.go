package util

import (
	"os"
	"strings"
)

func ExpandHome(path *string) {
	derefPath := *path
	if strings.Contains(derefPath, "~") {
		derefPath = strings.ReplaceAll(derefPath, "~", os.Getenv("HOME"))
	}
	*path = derefPath
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
