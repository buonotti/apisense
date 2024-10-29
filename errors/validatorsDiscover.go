package errors

import "github.com/joomcode/errorx"

var (
	ValidatorDiscoveryErrors = errorx.NewNamespace("validator_discovery")
	GoDiscoveryErrors        = ValidatorDiscoveryErrors.NewSubNamespace("go")
	RustDiscoveryErrors      = ValidatorDiscoveryErrors.NewSubNamespace("rust")
)

var (
	CannotDetermineValidatorPathError = ValidatorDiscoveryErrors.NewType("cannot_determine_validator_path")
	GoModFileEmptyError               = GoDiscoveryErrors.NewType("go_mod_empty")
	ModuleLineMalformedError          = GoDiscoveryErrors.NewType("module_line_malformed")
	CannotUnmarshalCargoTomlError     = RustDiscoveryErrors.NewType("cannot_unmarshal_cargotoml")
)
