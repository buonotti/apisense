package validators

import (
	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/validation"
)

func NewSchemaValidator() validation.Validator {
	return schemaValidator{}
}

type schemaValidator struct {
}

func (v schemaValidator) Name() string {
	return "schema"
}

func (v schemaValidator) Validate(item validation.Item) error {
	return validateSchema(item.Entries, item.Data)
}

func validateSchema(definitions []validation.SchemaEntry, data map[string]any) error {
	for _, definition := range definitions {
		value := data[definition.Name]
		if value == nil && definition.IsRequired {
			return errors.ValidationError.New("validation failed for field %s: field is required", definition.Name)
		}
		switch value.(type) {
		case string:
			if definition.Type != "string" {
				return errors.ValidationError.New("validation failed for field %s: expected type %s, got string", definition.Name, definition.Type)
			}
		case int:
			if definition.Type != "int" {
				return errors.ValidationError.New("validation failed for field %s: expected type %s, got int", definition.Name, definition.Type)
			}
		case float64:
			if definition.Type != "float" && definition.Type != "integer" {
				return errors.ValidationError.New("validation failed for field %s: expected type %s, got float", definition.Name, definition.Type)
			}
		case bool:
			if definition.Type != "bool" {
				return errors.ValidationError.New("validation failed for field %s: expected type %s, got bool", definition.Name, definition.Type)
			}
		case []any:
			if definition.Type != "array" {
				return errors.ValidationError.New("validation failed for field %s: expected type %s, got array", definition.Name, definition.Type)
			}
			arr, _ := value.([]any)
			if len(arr) == 0 && definition.IsRequired {
				return errors.ValidationError.New("validation failed for field %s: field is required (array empty)", definition.Name)
			}
			for _, item := range arr {
				err := validateSchema(definition.ChildEntries, item.(map[string]any))
				if err != nil {
					return err
				}
			}
		case map[string]any:
			if definition.Type != "object" {
				return errors.ValidationError.New("validation failed for field %s: expected type %s, got object", definition.Name, definition.Type)
			}
			err := validateSchema(definition.ChildEntries, value.(map[string]any))
			if err != nil {
				return err
			}
		default:
			return errors.ValidationError.New("validation failed for field %s: unknown type", definition.Name)
		}
	}
	return nil
}
