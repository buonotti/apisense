package errors

import (
	"os"

	"github.com/joomcode/errorx"

	"github.com/buonotti/odh-data-monitor/log"
)

var fatalTrait = errorx.RegisterTrait("fatal")

func HandleError(err error) {
	if err != nil {
		log.DefaultLogger.Error(err.Error())
		// fmt.Printf("Fatal error: %+v", err)
		os.Exit(1)
	}
}
