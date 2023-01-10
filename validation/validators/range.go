package validators

import (
	"strconv"

	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/validation"
)

// NewRangeValidator creates a new instance of a range validator
func NewRangeValidator() validation.Validator {
	return rangeValidator{}
}

// rangeValidator is a validator that validates the range of a field
type rangeValidator struct {
}

// Name returns the name of the validator: range
func (v rangeValidator) Name() string {
	return "range"
}

// Validate validates the given item by checking the range values in the
// definition, and return nil on success or an error on failure
func (v rangeValidator) Validate(item validation.PipelineItem) error {
	// go through all the definitions
	for _, schemaEntry := range item.SchemaEntries {
		// get the response value for the current schema schemaEntry
		value := item.Data[schemaEntry.Name]
		if value == nil {
			continue
		}

		// check the type of the value if the value is not a float or an int, skip the
		// check because the range validator doesn't apply to other types
		switch value.(type) {
		case int:
			// get the config values for min and max and parse them, then set them to +-inf
			// if they are none then compare the value
			valueInt, _ := value.(int64)
			minStr, _ := schemaEntry.Minimum.(string)
			maxStr, _ := schemaEntry.Maximum.(string)
			if minStr == "" || minStr == "none" {
				minStr = "-inf"
			}

			if maxStr == "" || maxStr == "none" {
				maxStr = "inf"
			}

			min, err := strconv.ParseInt(minStr, 10, 32)
			if err != nil {
				return errors.ValidationError.New("validation failed for field %s: cannot parse minimum value %s", schemaEntry.Name, schemaEntry.Minimum)
			}

			max, err := strconv.ParseInt(maxStr, 10, 32)
			if err != nil {
				return errors.ValidationError.New("validation failed for field %s: cannot parse maximum value %s", schemaEntry.Name, schemaEntry.Maximum)
			}

			if valueInt < min || valueInt > max {
				return errors.ValidationError.New("validation failed for field %s: expected value between %d and %d, got %d", schemaEntry.Name, min, max, value)
			}
		case float64:
			// do the same thing as above but for float64
			valueFloat, _ := value.(float64)
			minStr, _ := schemaEntry.Minimum.(string)
			maxStr, _ := schemaEntry.Maximum.(string)
			if minStr == "" || minStr == "none" {
				minStr = "-inf"
			}

			if maxStr == "" || maxStr == "none" {
				maxStr = "inf"
			}

			min, err := strconv.ParseFloat(minStr, 64)
			if err != nil {
				return errors.ValidationError.New("validation failed for field %s: cannot parse minimum value %s", schemaEntry.Name, schemaEntry.Minimum)
			}

			max, err := strconv.ParseFloat(maxStr, 64)
			if err != nil {
				return errors.ValidationError.New("validation failed for field %s: cannot parse maximum value %s", schemaEntry.Name, schemaEntry.Maximum)
			}

			if valueFloat < min || valueFloat > max {
				return errors.ValidationError.New("validation failed for field %s: expected value between %d and %d, got %d", schemaEntry.Name, min, max, value)
			}
		}
	}
	return nil
}

func (v rangeValidator) Fatal() bool {
	return true
}
