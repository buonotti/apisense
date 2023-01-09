package log

import (
	"github.com/sirupsen/logrus"
)

var DaemonLogger = logrus.WithFields(logrus.Fields{"system": "daemon"})
var DefaultLogger = logrus.WithFields(logrus.Fields{"system": "default"})
