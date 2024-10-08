package errors

import (
	"github.com/buonotti/apisense/log"
)

// CheckErr is a helper function that handles a given error. If the error is
// not nil it logs the error and exits with code 1. This function should only be
// used in the top level scope of the program or when an error cant be returned
func CheckErr(err error) {
	if err != nil {
		log.DefaultLogger().Fatal(err.Error())
	}
}
