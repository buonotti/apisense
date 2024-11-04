package validators

import (
	"github.com/buonotti/apisense/v2/errors"
	"github.com/santhosh-tekuri/jsonschema/v6"
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
	compiler := jsonschema.NewCompiler()
	err := compiler.AddResource("schema.json", item.Definition().ResponseSchema)
	if err != nil {
		return errors.ValidationError.Wrap(err, "add resource to schema failed")
	}
	schema, err := compiler.Compile("schema.json")
	if err != nil {
		return errors.ValidationError.Wrap(err, "validation of schema failed")
	}
	err = schema.Validate(item.Response().RawData)
	if err != nil {
		return errors.ValidationError.Wrap(err, "response schema is invalid")
	}

	return nil
}

func (v schemaValidator) IsFatal() bool {
	return true
}

func (_ schemaValidator) IsSlim() bool {
	return false
}
