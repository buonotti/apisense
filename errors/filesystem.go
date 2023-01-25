package errors

import (
	"github.com/joomcode/errorx"
)

// FileSystemErrors is the namespace holding all file system related errors
var FileSystemErrors = errorx.NewNamespace("filesystem")

// FileNotFoundError is returned when a file is not found
var FileNotFoundError = FileSystemErrors.NewType("file_not_found", fatalTrait)

// CannotParseDefinitionFileError is returned when the daemon fails to parse the definition file
var CannotParseDefinitionFileError = FileSystemErrors.NewType("cannot_parse_definition_file", fatalTrait)

// CannotWriteFileError is returned when a functions fails to read a file
var CannotWriteFileError = FileSystemErrors.NewType("cannot_write_file", fatalTrait)

// CannotReadFileError is returned when a functions fails to read a file
var CannotReadFileError = FileSystemErrors.NewType("cannot_read_file", fatalTrait)

// CannotReadLockFileError is returned when the lock file cannot be read or does not exist
var CannotReadLockFileError = FileSystemErrors.NewType("cannot_read_lock_file", fatalTrait)

// CannotLockFileError is returned when the lock file cannot be locked by the current process. It generally means that another daemon is already running.
var CannotLockFileError = FileSystemErrors.NewType("cannot_lock_file", fatalTrait)

// CannotUnlockFileError is returned when the already locked lock file cannot be unlocked
var CannotUnlockFileError = FileSystemErrors.NewType("cannot_unlock_file", fatalTrait)

// CannotCreateDirectoryError is returned when any of the setup functions fails to create the required directories
var CannotCreateDirectoryError = FileSystemErrors.NewType("cannot_create_directory", fatalTrait)

var CannotFindReportFile = FileSystemErrors.NewType("cannot_find_report_file", fatalTrait)
