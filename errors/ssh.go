package errors

import (
	"github.com/joomcode/errorx"
)

var SSHErrors = errorx.NewNamespace("ssh")
var CannotCreateSSHServerError = SSHErrors.NewType("cannot_create_ssh_server", fatalTrait)
var CannotStopSSHServerError = SSHErrors.NewType("cannot_stop_ssh_server")
