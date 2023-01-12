package conversion

import (
	"encoding/json"

	"github.com/buonotti/odh-data-monitor/validation"
)

func Json() Converter {
	return jsonConverter{}
}

type jsonConverter struct{}

func (jsonConverter) Convert(report validation.Report) ([]byte, error) {
	d, err := json.Marshal(report)
	return d, err
}
