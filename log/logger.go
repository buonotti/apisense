package log

import (
	"github.com/sirupsen/logrus"
)

var DaemonLogger = logrus.WithFields(logrus.Fields{"system": "daemon"})
