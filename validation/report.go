package validation

import (
	"encoding/json"
	"fmt"
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
	Time    ReportTime          // Time is the timestamp of the report
	Results []ValidatedEndpoint // Results is a collection of ValidatedEndpoint holding the validation results
}

type ReportTime time.Time

//goland:noinspection GoMixedReceiverTypes
func (t ReportTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02T15:04:05.000Z"))), nil
}

//goland:noinspection GoMixedReceiverTypes
func (t *ReportTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	tt, err := time.Parse("2006-01-02T15:04:05.000Z", s)
	if err != nil {
		return err
	}
	*t = ReportTime(tt)
	return nil
}

//goland:noinspection GoMixedReceiverTypes
func (t ReportTime) String() string {
	return time.Time(t).Format("2006-01-02T15:04:05.000Z")
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
func RawReports() ([]map[string]any, error) {
	files, err := os.ReadDir(ReportLocation())
	errors.HandleError(err)

	reports := make([]map[string]any, 0)
	for _, file := range files {
		if !file.IsDir() {
			fileName := ReportLocation() + "/" + file.Name()
			content, err := os.ReadFile(fileName)
			if err != nil {
				return nil, errors.CannotReadFileError.Wrap(err, "cannot read file:"+fileName)
			}
			item := make(map[string]any)
			err = json.Unmarshal(content, &item)
			if err != nil {
				return nil, errors.CannotUnmarshalReportFileError.Wrap(err, "cannot unmarshal file:"+fileName)
			}
			reports = append(reports, item)
		}
	}
	return reports, nil
}
