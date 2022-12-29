package errors

import (
	"github.com/joomcode/errorx"
)

var FileSystemErrors = errorx.NewNamespace("filesystem")
var FileNotFound = FileSystemErrors.NewType("file_not_found", fatalTrait)
var CannotParseDefinitionFile = FileSystemErrors.NewType("cannot_parse_definition_file", fatalTrait)


