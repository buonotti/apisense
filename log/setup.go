package log

import (
	"os"

	"github.com/apex/log"
	"github.com/spf13/viper"
)

func hasLogFile() bool {
	return logFile != nil
}

var logFile *os.File = nil

func CloseLogFile() error {
	if logFile != nil {
		err := logFile.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// Setup reads the config and configures log level and log output of all loggers
func Setup() error {
	logFilePath := viper.GetString("log.file")
	if logFilePath != "" {
		osFile, err := os.OpenFile(viper.GetString("log.file"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		logFile = osFile
	}

	log.SetHandler(newHandler())

	log.SetLevelFromString(viper.GetString("log.level"))

	return nil
}
