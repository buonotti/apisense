package errors

import (
	"github.com/joomcode/errorx"
)

var (
	ValidationErrors                 = errorx.NewNamespace("validation")
	ValidationError                  = ValidationErrors.NewType("validation_error")
	CannotSerializeItemError         = ValidationErrors.NewType("cannot_serialize_item")
	ExternalValidatorParseError      = ValidationErrors.NewType("external_validator_parse_error")
	VariableValueLengthMismatchError = ValidationErrors.NewType("variable_value_length_mismatch")
	CannotUnmarshalReportFileError   = ValidationErrors.NewType("cannot_unmarshal_report_file")
	InvalidApiResponseError          = ValidationErrors.NewType("invalid_api_response")
)
