package errors

import (
	"github.com/joomcode/errorx"
)

var ValidationErrors = errorx.NewNamespace("validation")
var ValidationError = ValidationErrors.NewType("validation_error")
var CannotSerializeItemError = ValidationErrors.NewType("cannot_serialize_item", fatalTrait)
var ExternalValidatorParseError = ValidationErrors.NewType("external_validator_parse_error", fatalTrait)
var VariableValueLengthMismatchError = ValidationErrors.NewType("variable_value_length_mismatch", fatalTrait)
var CannotUnmarshalReportFileError = ValidationErrors.NewType("cannot_unmarshal_report_file", fatalTrait)
