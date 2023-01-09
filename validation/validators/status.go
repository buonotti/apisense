package validators

import (
	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/validation"
)

func NewStatusValidator() validation.Validator {
	return statusValidator{
		OkStatus: 200,
	}
}

func NewStatusValidatorC(code int) validation.Validator {
	return statusValidator{
		OkStatus: code,
	}
}

type statusValidator struct {
	OkStatus int
}

func (v statusValidator) Name() string {
	return "status"
}

func (v statusValidator) Validate(item validation.PipelineItem) error {
	if item.Code != v.OkStatus {
		return errors.ValidationError.New("validation failed for endpoint %s: expected status code 200, got %d", item.Endpoint, item.Code)
	}
	return nil
}
