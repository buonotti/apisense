package errors

import (
	"github.com/joomcode/errorx"
)

var (
	ValidationErrors                 = errorx.NewNamespace("validation")
	ValidationError                  = ValidationErrors.NewType("validation_error")
	CannotSerializeItemError         = ValidationErrors.NewType("cannot_serialize_item", fatalTrait)
	ExternalValidatorParseError      = ValidationErrors.NewType("external_validator_parse_error", fatalTrait)
	VariableValueLengthMismatchError = ValidationErrors.NewType("variable_value_length_mismatch", fatalTrait)
	CannotUnmarshalReportFileError   = ValidationErrors.NewType("cannot_unmarshal_report_file", fatalTrait)
)
