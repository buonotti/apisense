package validation

import (
	"reflect"
	"strings"
	"time"
)

type datetimeValidator struct{}

func (v datetimeValidator) Validate(field string, value any, arg string) error {
	val := reflect.ValueOf(value)
	if val.IsZero() {
		return nil
	}
	if val.Kind() != reflect.String {
		return errorf(field, "datetime", "", "value is not a string")
	}
	if arg == "" {
		return errorf(field, "datetime", val.String(), "missing format string")
	}

	_, err := time.Parse(arg, val.String())
	if err != nil {
		errStr := strings.ReplaceAll(err.Error(), "\"", "'")
		return errorf(field, "datetime", val.String(), errStr)
	}
	return nil
}
