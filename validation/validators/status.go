package validators

import (
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/validation/fetcher"
)

// NewStatusValidator returns a new status validator that checks for status codes other than 200
func NewStatusValidator() Validator {
	return statusValidator{
		OkStatus: 200,
	}
}

// NewStatusValidatorC returns a new status validator that checks for status codes other than the given code
func NewStatusValidatorC(code int) Validator {
	return statusValidator{
		OkStatus: code,
	}
}

// statusValidator is a validator that checks if the item has the return code matching some given code
type statusValidator struct {
	OkStatus int // OkStatus is the status code that is allowed to get returned. Reports other codes will result in a failure
}

// Name returns the name of the validator: status
func (v statusValidator) Name() string {
	return "status"
}

// Validate checks for each item if the status code of the response matches the given status code
func (v statusValidator) Validate(item fetcher.TestCase) error {
	if item.Code != v.OkStatus {
		return errors.ValidationError.New("validation failed for endpoint %s: expected status code 200, got %d", item.Url, item.Code)
	}
	return nil
}

func (v statusValidator) IsFatal() bool {
	return true
}
