package validators

import (
	"strconv"

	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/validation"
)

func NewRangeValidator() validation.Validator {
	return rangeValidator{}
}

type rangeValidator struct {
}

func (v rangeValidator) Name() string {
	return "range"
}

func (v rangeValidator) Validate(item validation.Item) error {
	return validateRange(item.Entries, item.Data)
}

func validateRange(definitions []validation.SchemaEntry, data map[string]any) error {
	for _, definition := range definitions {
		value := data[definition.Name]
		if value == nil {
			continue
		}
		switch value.(type) {
		case int:
			valueInt, _ := value.(int64)
			minStr, _ := definition.Minimum.(string)
			maxStr, _ := definition.Maximum.(string)
			if minStr == "" || minStr == "none" {
				minStr = "-inf"
			}
			if maxStr == "" || maxStr == "none" {
				maxStr = "inf"
			}
			min, err := strconv.ParseInt(minStr, 10, 32)
			if err != nil {
				return errors.ValidationError.New("validation failed for field %s: cannot parse minimum value %s", definition.Name, definition.Minimum)
			}
			max, err := strconv.ParseInt(maxStr, 10, 32)
			if err != nil {
				return errors.ValidationError.New("validation failed for field %s: cannot parse maximum value %s", definition.Name, definition.Maximum)
			}
			if valueInt < min || valueInt > max {
				return errors.ValidationError.New("validation failed for field %s: expected value between %d and %d, got %d", definition.Name, min, max, value)
			}
		case float64:
			valueFloat, _ := value.(float64)
			minStr, _ := definition.Minimum.(string)
			maxStr, _ := definition.Maximum.(string)
			if minStr == "" || minStr == "none" {
				minStr = "-inf"
			}
			if maxStr == "" || maxStr == "none" {
				maxStr = "inf"
			}
			min, err := strconv.ParseFloat(minStr, 64)
			if err != nil {
				return errors.ValidationError.New("validation failed for field %s: cannot parse minimum value %s", definition.Name, definition.Minimum)
			}
			max, err := strconv.ParseFloat(maxStr, 64)
			if err != nil {
				return errors.ValidationError.New("validation failed for field %s: cannot parse maximum value %s", definition.Name, definition.Maximum)
			}
			if valueFloat < min || valueFloat > max {
				return errors.ValidationError.New("validation failed for field %s: expected value between %d and %d, got %d", definition.Name, min, max, value)
			}
		}
	}
	return nil
}
