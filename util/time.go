package util

import (
	"fmt"
	"strings"
	"time"
)

// ApisenseTime is the time format used in everything apisense
type ApisenseTime time.Time

// ApisenseTimeFormat is the go time format representation of ApisenseTime
const ApisenseTimeFormat = "2006-01-02T15-04-05.000Z"

func (t ApisenseTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(t).Format(ApisenseTimeFormat))), nil
}

func (t *ApisenseTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	tt, err := time.Parse(ApisenseTimeFormat, s)
	if err != nil {
		return err
	}
	*t = ApisenseTime(tt)
	return nil
}

// String returns the formatted representation according to ApisenseTimeFormat
func (t ApisenseTime) String() string {
	return time.Time(t).Format(ApisenseTimeFormat)
}

// ParseTime parses a string into an ApisenseTime according to ApisenseTimeFormat
func ParseTime(s string) (ApisenseTime, error) {
	parsedTime, err := time.Parse(ApisenseTimeFormat, s)
	return ApisenseTime(parsedTime), err
}
