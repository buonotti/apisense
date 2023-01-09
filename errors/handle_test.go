package errors_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/buonotti/odh-data-monitor/errors"
)

func TestHandleError_OnError(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		errors.HandleError(fmt.Errorf("test error"))
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestHandleError_OnError")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestHandleError_OnNil(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		errors.HandleError(nil)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestHandleError_OnNil")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		t.Fatalf("process ran with err %v, want exit status 0", err)
	}
	return
}
