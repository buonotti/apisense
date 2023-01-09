package log

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Setup() error {
	doPrettyLog := viper.GetBool("daemon.log.pretty")
	forceColorLog := viper.GetBool("daemon.log.force-color")
	logLevel := viper.GetString("daemon.log.level")
	logFileName := viper.GetString("daemon.log.file")

	if doPrettyLog {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:  forceColorLog,
			PadLevelText: true,
		})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	logrus.SetLevel(lvl)

	if logFileName != "" {
		logFile, err := os.Open(logFileName)
		if err != nil {
			return err
		}
		logrus.SetOutput(logFile)
	}

	return nil
}
