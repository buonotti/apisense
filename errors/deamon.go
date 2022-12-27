package errors

import (
	"github.com/joomcode/errorx"
)

var DaemonErrors = errorx.NewNamespace("daemon")
var CannotReadStatusError = DaemonErrors.NewType("cannot_read_status")
var CannotReadLockFileError = DaemonErrors.NewType("cannot_read_lock_file", fatalTrait)
var CannotLockFileError = DaemonErrors.NewType("cannot_lock_file")
var CannotUnlockFileError = DaemonErrors.NewType("cannot_unlock_file")
var CannotCreateDirectoryError = DaemonErrors.NewType("cannot_create_directory", fatalTrait)
