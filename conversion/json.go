package conversion

import (
	"encoding/json"

	"github.com/buonotti/apisense/validation"
)

func Json() Converter {
	return jsonConverter{}
}

type jsonConverter struct{}

func (jsonConverter) Convert(reports ...validation.Report) ([]byte, error) {
	if len(reports) == 1 {
		return json.Marshal(reports[0])
	} else {
		return json.Marshal(reports)
	}
}
