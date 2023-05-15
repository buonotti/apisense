package validation

import (
	"fmt"
	"reflect"
	"strings"
)

type Failure error

func errorf(field string, validator string, value string, details string) error {
	base := fmt.Sprintf("field %s failed validation %s", field, validator)
	if value != "" {
		base += fmt.Sprintf(" with value %s", value)
	}
	if details != "" {
		base += fmt.Sprintf(" (%s)", details)
	}
	return fmt.Errorf(base)
}

func Validate(obj any) error {
	t := reflect.TypeOf(obj)

	accessAsPointer := false

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		accessAsPointer = true
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		validatorTag := field.Tag.Get("validate")
		validatorSpecs := strings.Split(validatorTag, ",")
		for _, validatorSpec := range validatorSpecs {
			if validatorSpec == "" {
				continue
			}
			splitSpec := strings.Split(validatorSpec, "=")
			validator := validators[splitSpec[0]]
			if validator == nil {
				return fmt.Errorf("validator %s not found", validator)
			}
			var value any
			if accessAsPointer {
				valuePtr := reflect.ValueOf(obj)
				value = reflect.Indirect(valuePtr).Field(i).Interface()
			} else {
				value = reflect.ValueOf(obj).Field(i).Interface()
			}
			arg := ""
			if len(splitSpec) > 1 {
				arg = splitSpec[1]
			}
			if err := validator.Validate(field.Name, value, arg); err != nil {
				return err
			}
		}
	}
	return nil
}
