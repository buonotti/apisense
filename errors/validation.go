package errors

import (
	"github.com/joomcode/errorx"
)

// ValidationErrors is the namespace holding all validation related errors
var ValidationErrors = errorx.NewNamespace("validation")

// ValidationError is the general error returned when a validator fails
var ValidationError = ValidationErrors.NewType("validation_error")

// CannotSerializeItemError is returned when the pipeline fails to serialize an item to pass it to an external validator
var CannotSerializeItemError = ValidationErrors.NewType("cannot_serialize_item", fatalTrait)

// ExternalValidatorParseError is returned when the external validator declared in the config has an invalid structure
var ExternalValidatorParseError = ValidationErrors.NewType("external_validator_parse_error", fatalTrait)

// VariableValueLengthMismatchError is returned when the number of values in the variables is not equal across all the variables
var VariableValueLengthMismatchError = ValidationErrors.NewType("variable_value_length_mismatch", fatalTrait)

// CannotUnmarshalReportFileError is returned when a report file cannot be deserialized
var CannotUnmarshalReportFileError = ValidationErrors.NewType("cannot_unmarshal_report_file", fatalTrait)
