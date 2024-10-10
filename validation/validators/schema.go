package validators

import (
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation/definitions"
	"github.com/xeipuuv/gojsonschema"
)

// NewSchemaValidator returns a new schema validator
func NewSchemaValidator() Validator {
	return schemaValidator{}
}

// schemaValidator is a validator that validates the response data against the result schema in the definition
type schemaValidator struct{}

// Name returns the name of the validator: schema
func (v schemaValidator) Name() string {
	return "schema"
}

// Validate validates the given items schema and return nil on success or an error on failure
func (v schemaValidator) Validate(item ValidationItem) error {
	schemaLoader := gojsonschema.NewGoLoader(item.Definition.ResponseSchema)
	documentLoader := gojsonschema.NewGoLoader(item.Response.RawData)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return errors.ValidationError.Wrap(err, "validation of schema failed")
	}
	if !result.Valid() {
		return errors.ValidationError.New("response has invalid schema: " + util.FirstOrDefault(util.Map(result.Errors(), func(e gojsonschema.ResultError) string {
			return e.String()
		}), "<?>"))
	}
	return nil
}

// validateSchema validates the result against the schema and return nil on success or an error on failure.
// The function recursively checks child object or array definitions
func validateSchema(schemaEntries []definitions.SchemaEntry, data map[string]any) error {
	for _, schemaEntry := range schemaEntries {
		// get the response value for the current schema schemaEntry
		value := data[schemaEntry.Name]

		// check if the value is nil and if it is required
		if value == nil && schemaEntry.IsRequired {
			return errors.ValidationError.New("validation failed for field %s: field is required", schemaEntry.Name)
		}

		// check the type of the value. for each type, check if the value matches the type in the definition
		switch value := value.(type) {
		case string:
			if schemaEntry.Type != "string" {
				return errors.ValidationError.New("validation failed for field %s: expected type %s, got string", schemaEntry.Name, schemaEntry.Type)
			}
		case int:
			if schemaEntry.Type != "number" && schemaEntry.Type != "integer" {
				return errors.ValidationError.New("validation failed for field %s: expected type %s, got int", schemaEntry.Name, schemaEntry.Type)
			}
		case float64:
			if schemaEntry.Type != "number" {
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
			arr := value

			// check if the array is empty and if it is required
			if len(arr) == 0 && schemaEntry.IsRequired {
				return errors.ValidationError.New("validation failed for field %s: field is required (array empty)", schemaEntry.Name)
			}

			// validate each item in the array by recursively calling this function on each
			// item in the array using the schema in the children property of the current
			// schema schemaEntry
			for _, item := range arr {
				err := validateSchema(schemaEntry.Fields, item.(map[string]any))
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
			err := validateSchema(schemaEntry.Fields, value)
			if err != nil {
				return err
			}
		default:
			return errors.ValidationError.New("validation failed for field %s: unknown type", schemaEntry.Name)
		}
	}
	return nil
}

func (v schemaValidator) IsFatal() bool {
	return true
}
