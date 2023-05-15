package validation

import (
	"reflect"
	"strings"
)

type oneofValidator struct{}

func (v oneofValidator) Validate(field string, value any, arg string) error {
	val := reflect.ValueOf(value)
	if val.IsZero() {
		return nil
	}
	if arg == "" {
		return errorf(field, "oneof", "", "missing arguments")
	}
	values := strings.Split(arg, " ")
	for _, v := range values {
		if val.String() == v {
			return nil
		}
	}
	return errorf(field, "oneof", val.String(), "value is not in the list of allowed values")
}
