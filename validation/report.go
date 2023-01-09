package validation

import (
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

func ReportLocation() string {
	return filepath.FromSlash(viper.GetString("daemon.reports-dir"))
}

type Report struct {
	Time    time.Time
	Results []ValidatedEndpoint
}
