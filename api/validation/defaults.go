package validation

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func FillDefaults(obj any) error {
	t := reflect.TypeOf(obj)

	accessAsPointer := false

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		accessAsPointer = true
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		defaultTag := field.Tag.Get("default")
		if defaultTag == "" {
			continue
		}
		var value any
		if accessAsPointer {
			valuePtr := reflect.ValueOf(obj)
			value = reflect.Indirect(valuePtr).Field(i).Interface()
		} else {
			value = reflect.ValueOf(obj).Field(i).Interface()
		}
		if reflect.ValueOf(value).IsZero() {
			var field reflect.Value
			if accessAsPointer {
				valuePtr := reflect.ValueOf(obj)
				field = reflect.Indirect(valuePtr).Field(i)
			} else {
				field = reflect.ValueOf(obj).Field(i)
			}
			parsedValue, err := parseValue(field, defaultTag)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(parsedValue))
		}
	}
	return nil
}

func parseValue(value reflect.Value, bindValue string) (any, error) {
	switch value.Kind() {
	case reflect.String:
		return bindValue, nil
	case reflect.Int:
		return strconv.Atoi(bindValue)
	case reflect.Bool:
		return strconv.ParseBool(bindValue)
	case reflect.Float64:
		return strconv.ParseFloat(bindValue, 64)
	case reflect.Float32:
		return strconv.ParseFloat(bindValue, 32)
	case reflect.Slice:
		stringValues := strings.Split(bindValue, " ")
		slice := reflect.MakeSlice(value.Type(), len(stringValues), len(stringValues))
		for i, stringValue := range stringValues {
			parsedValue, err := parseValue(slice.Index(i), stringValue)
			if err != nil {
				return nil, err
			}
			slice.Index(i).Set(reflect.ValueOf(parsedValue))
		}
	case reflect.Struct:
		var val any
		err := json.Unmarshal([]byte(bindValue), &val)
		if err != nil {
			return nil, err
		}
		return val, nil
	}
	return nil, fmt.Errorf("unsupported type %s", value.Kind())
}
