package errors

import (
	"fmt"
	"os"

	"github.com/joomcode/errorx"
)

var fatalTrait = errorx.RegisterTrait("fatal")

func HandleError(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %+v", err)
		os.Exit(1)
	}
}
