package errors

import (
	"github.com/joomcode/errorx"
)

// SSHErrors is the namespace holding all SSH related errors
var SSHErrors = errorx.NewNamespace("ssh")

// CannotCreateSSHServerError is returned when the function wish.NewServer fails to create a new ssh server
var CannotCreateSSHServerError = SSHErrors.NewType("cannot_create_ssh_server", fatalTrait)

// CannotStopSSHServerError is returned when the ssh server fails to stop
var CannotStopSSHServerError = SSHErrors.NewType("cannot_stop_ssh_server")
