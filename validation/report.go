package validation

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func ReportLocation() string {
	path := viper.GetString("daemon.reports-dir")
	if strings.Contains(path, "~") {
		path = strings.Replace(path, "~", os.Getenv("HOME"), 1)
	}
	return filepath.FromSlash(path)
}

type Report struct {
	Time    time.Time
	Results []ValidatedEndpoint
}
