package log

import (
	"github.com/apex/log"
)

var DaemonLogger = log.WithFields(log.Fields{"system": "daemon"})
var DefaultLogger = log.WithFields(log.Fields{"system": "default"})
var SSHLogger = log.WithFields(log.Fields{"system": "ssh"})
var ApiLogger = log.WithFields(log.Fields{"system": "api"})
var CliLogger = log.WithFields(log.Fields{"system": "cli"})
