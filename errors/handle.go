package errors

import (
	"fmt"
	"os"

	"github.com/joomcode/errorx"

	"github.com/buonotti/odh-data-monitor/log"
)

// fatalTrait is a trait that can be added to any error type to make it fatal and make the program exit with code 1
var fatalTrait = errorx.RegisterTrait("fatal")

// HandleError is a helper function that handles a given error. If the error is
// not nil it logs the error and exits with code 1. This function should only be
// used in the top level scope of the program or when an error cant be returned
func HandleError(err error) {
	if _, ok := err.(*errorx.Error); ok {
		handleErrorxError(err.(*errorx.Error))
	} else {
		handleError(err)
	}
}

func handleErrorxError(err *errorx.Error) {
	if err != nil && err.Unwrap() != nil {
		fmt.Printf("%+v", err) // TODO
		if err.HasTrait(fatalTrait) {
			log.DefaultLogger.Error(err.Error())
			os.Exit(1)
		} else {
			log.DefaultLogger.Warn(err.Error())
		}
	}
}

func handleError(err error) {
	if err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}
