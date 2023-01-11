package validation

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"

	"github.com/buonotti/odh-data-monitor/errors"
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
	Id      string              // Id is a unique identifier for each report
	Time    time.Time           // Time is the timestamp of the report
	Results []ValidatedEndpoint // Results is a collection of ValidatedEndpoint holding the validation results
}

// Reports returns all the reports in the report directory
func Reports() ([]Report, error) {
	files, err := os.ReadDir(ReportLocation())
	errors.HandleError(err)

	reports := make([]Report, 0)
	for _, file := range files {
		if !file.IsDir() {
			fileName := ReportLocation() + "/" + file.Name()
			content, err := os.ReadFile(fileName)
			if err != nil {
				return nil, errors.CannotReadFileError.Wrap(err, "cannot read file:"+fileName)
			}

			var report Report
			err = json.Unmarshal(content, &report)
			if err != nil {
				return nil, errors.CannotUnmarshalReportFileError.Wrap(err, "cannot unmarshal file:"+fileName)
			}

			reports = append(reports, report)
		}
	}
	return reports, nil
}

// RawReports return all the reports in the report directory without unmarshalling them
func RawReports() ([]string, error) {
	files, err := os.ReadDir(ReportLocation())
	errors.HandleError(err)

	reports := make([]string, 0)
	for _, file := range files {
		if !file.IsDir() {
			fileName := ReportLocation() + "/" + file.Name()
			content, err := os.ReadFile(fileName)
			if err != nil {
				return nil, errors.CannotReadFileError.Wrap(err, "cannot read file:"+fileName)
			}
			reports = append(reports, string(content))
		}
	}
	return reports, nil
}
