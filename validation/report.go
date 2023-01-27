package validation

import (
	"encoding/json"
	"os"

	"github.com/buonotti/apisense/errors"
)

// ReportLocation returns the output directory where all reports are stored
func ReportLocation() string {
	return os.Getenv("HOME") + "/apisense/reports"
}

// Report is a report of a test run
type Report struct {
	Id        string              `json:"id"`        // Id is a unique identifier for each report
	Time      ReportTime          `json:"time"`      // Time is the timestamp of the report
	Endpoints []ValidatedEndpoint `json:"endpoints"` // Endpoints is a collection of ValidatedEndpoint holding the validation results
}

func GetReport(filename string) (*Report, error) {
	files, err := os.ReadDir(ReportLocation())
	errors.HandleError(err)

	for _, file := range files {
		if !file.IsDir() && file.Name() == filename {
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

			return &report, nil
		}
	}
	return nil, errors.CannotFindReportFile.New("cannot find report file: " + filename)
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
