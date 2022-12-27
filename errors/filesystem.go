package errors

import (
	"github.com/joomcode/errorx"
)

var FileSystemErrors = errorx.NewNamespace("filesystem")
var FileNotFound = FileSystemErrors.NewType("file_not_found", fatalTrait)


