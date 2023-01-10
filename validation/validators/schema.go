package validators

import (
	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/validation"
)

// NewSchemaValidator returns a new schema validator
func NewSchemaValidator() validation.Validator {
	return schemaValidator{}
}

// schemaValidator is a validator that validates the response data against the result schema in the definition
type schemaValidator struct {
}

// Name returns the name of the validator: schema
func (v schemaValidator) Name() string {
	return "schema"
}

// Validate validates the given items schema and return nil on success or an error on failure
func (v schemaValidator) Validate(item validation.PipelineItem) error {
	return validateSchema(item.SchemaEntries, item.Data)
}

// validateSchema validates the result against the schema and return nil on success or an error on failure.
// The function recursively checks child object or array definitions
func validateSchema(schemaEntries []validation.SchemaEntry, data map[string]any) error {
	for _, schemaEntry := range schemaEntries {
		// get the response value for the current schema schemaEntry
		value := data[schemaEntry.Name]

		// check if the value is nil and if it is required
		if value == nil && schemaEntry.IsRequired {
			return errors.ValidationError.New("validation failed for field %s: field is required", schemaEntry.Name)
		}

		// check the type of the value. for each type, check if the value matches the type in the definition
		switch value.(type) {
		case string:
			if schemaEntry.Type != "string" {
				return errors.ValidationError.New("validation failed for field %s: expected type %s, got string", schemaEntry.Name, schemaEntry.Type)
			}
		case int:
			if schemaEntry.Type != "int" {
				return errors.ValidationError.New("validation failed for field %s: expected type %s, got int", schemaEntry.Name, schemaEntry.Type)
			}
		case float64:
			if schemaEntry.Type != "float" && schemaEntry.Type != "integer" {
				return errors.ValidationError.New("validation failed for field %s: expected type %s, got float", schemaEntry.Name, schemaEntry.Type)
			}
		case bool:
			if schemaEntry.Type != "bool" {
				return errors.ValidationError.New("validation failed for field %s: expected type %s, got bool", schemaEntry.Name, schemaEntry.Type)
			}
		case []any:
			if schemaEntry.Type != "array" {
				return errors.ValidationError.New("validation failed for field %s: expected type %s, got array", schemaEntry.Name, schemaEntry.Type)
			}

			// if the value is an array cast it
			arr, _ := value.([]any)

			// check if the array is empty and if it is required
			if len(arr) == 0 && schemaEntry.IsRequired {
				return errors.ValidationError.New("validation failed for field %s: field is required (array empty)", schemaEntry.Name)
			}

			// validate each item in the array by recursively calling this function on each
			// item in the array using the schema in the children property of the current
			// schema schemaEntry
			for _, item := range arr {
				err := validateSchema(schemaEntry.ChildEntries, item.(map[string]any))
				if err != nil {
					return err
				}
			}
		case map[string]any:
			if schemaEntry.Type != "object" {
				return errors.ValidationError.New("validation failed for field %s: expected type %s, got object", schemaEntry.Name, schemaEntry.Type)
			}

			// do the similar thing as above for arrays but for objects (that means we do the
			// same thing, we just check only one value instead of an array)
			err := validateSchema(schemaEntry.ChildEntries, value.(map[string]any))
			if err != nil {
				return err
			}
		default:
			return errors.ValidationError.New("validation failed for field %s: unknown type", schemaEntry.Name)
		}
	}
	return nil
}

func (v schemaValidator) Fatal() bool {
	return true
}
