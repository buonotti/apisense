package conversion

import (
	"encoding/json"

	"github.com/buonotti/odh-data-monitor/validation"
)

func Json() Converter {
	return jsonConverter{}
}

type jsonConverter struct{}

func (jsonConverter) Convert(reports ...validation.Report) ([]byte, error) {
	d, err := json.Marshal(reports)
	return d, err
}
