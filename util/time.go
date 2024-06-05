package util

import (
	"fmt"
	"strings"
	"time"
)

type ApisenseTime time.Time

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

func (t ApisenseTime) String() string {
	return time.Time(t).Format(ApisenseTimeFormat)
}

func ParseTime(s string) (ApisenseTime, error) {
	time, err := time.Parse(ApisenseTimeFormat, s)
	return ApisenseTime(time), err
}
