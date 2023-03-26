package pipeline

import (
	"fmt"
	"strings"
	"time"
)

type ReportTime time.Time

const ReportTimeFormat = "2006-01-02T15-04-05.000Z"

//goland:noinspection GoMixedReceiverTypes
func (t ReportTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(t).Format(ReportTimeFormat))), nil
}

//goland:noinspection GoMixedReceiverTypes
func (t *ReportTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	tt, err := time.Parse(ReportTimeFormat, s)
	if err != nil {
		return err
	}
	*t = ReportTime(tt)
	return nil
}

//goland:noinspection GoMixedReceiverTypes
func (t ReportTime) String() string {
	return time.Time(t).Format(ReportTimeFormat)
}
