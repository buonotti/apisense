package validation

import (
	"github.com/goccy/go-reflect"
)

type requiredValidator struct{}

func (v requiredValidator) Validate(field string, value any, _ string) error {
	if reflect.ValueOf(value).IsZero() {
		return errorf(field, "required", "", "")
	}
	return nil
}
