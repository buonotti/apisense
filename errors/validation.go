package errors

import (
	"github.com/joomcode/errorx"
)

var (
	ValidationErrors                 = errorx.NewNamespace("validation")
	InvalidDefinitionErrors          = ValidationErrors.NewSubNamespace("definition")
	DuplicateDefinitionError         = InvalidDefinitionErrors.NewType("duplicate_definitions")
	NoBaseUrlError                   = InvalidDefinitionErrors.NewType("no_base_url")
	VariableValueLengthMismatchError = InvalidDefinitionErrors.NewType("variable_value_length_mismatch")
	InvalidFormatError               = InvalidDefinitionErrors.NewType("invalid_format")
	TestCaseNamesLengthMismatchError = InvalidDefinitionErrors.NewType("test_case_names_length_mismatch")
	InvalidSchemaError               = InvalidDefinitionErrors.NewType("invalid_schema")
	InvalidCharacterError            = InvalidDefinitionErrors.NewType("invalid_character")
	ValidationError                  = ValidationErrors.NewType("validation_error")
	CannotSerializeItemError         = ValidationErrors.NewType("cannot_serialize_item")
	ExternalValidatorParseError      = ValidationErrors.NewType("external_validator_parse_error")
	CannotUnmarshalReportFileError   = ValidationErrors.NewType("cannot_unmarshal_report_file")
	InvalidApiResponseError          = ValidationErrors.NewType("invalid_api_response")
)
