package validation

import (
	"fmt"
	"strings"
	"time"
)

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
