package validation

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// ReportLocation returns the output directory where all reports are stored
func ReportLocation() string {
	path := viper.GetString("daemon.reports-dir")
	if strings.Contains(path, "~") {
		path = strings.Replace(path, "~", os.Getenv("HOME"), 1)
	}
	return filepath.FromSlash(path)
}

// Report is a report of a test run
type Report struct {
	Time    time.Time           // Time is the timestamp of the report
	Results []ValidatedEndpoint // Results is a collection of ValidatedEndpoint holding the validation results
}
