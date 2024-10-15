package conversion

import (
	"github.com/buonotti/apisense/validation/pipeline"
	"github.com/goccy/go-json"
)

func Json() Converter {
	return jsonConverter{}
}

type jsonConverter struct{}

func (jsonConverter) Convert(reports ...pipeline.Report) ([]byte, error) {
	if len(reports) == 1 {
		return json.Marshal(reports[0])
	} else {
		return json.Marshal(reports)
	}
}
