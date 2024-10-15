package validation

// Validator represents something capable of validating a struct field
type Validator interface {
	Validate(field string, value any, arg string) error
}

var validators = map[string]Validator{
	"required": requiredValidator{},
	"datetime": datetimeValidator{},
	"oneof":    oneofValidator{},
}
