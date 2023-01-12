package conversion

import (
	"github.com/buonotti/odh-data-monitor/validation"
)

type Converter interface {
	Convert(report validation.Report) ([]byte, error)
}
