package conversion

import (
	"encoding/xml"

	"github.com/buonotti/apisense/validation"
)

func Xml() Converter {
	return xmlConverter{}
}

type xmlConverter struct{}

type response struct {
	Reports []validation.Report `xml:"Report"`
}

func (xmlConverter) Convert(reports ...validation.Report) ([]byte, error) {
	if len(reports) == 1 {
		return xml.Marshal(reports[0])
	} else {
		return xml.Marshal(response{reports})
	}
}
