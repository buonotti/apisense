package conversion

import (
	"encoding/xml"

	"github.com/buonotti/odh-data-monitor/validation"
)

func Xml() Converter {
	return xmlConverter{}
}

type xmlConverter struct{}

type response struct {
	Reports []validation.Report `xml:"Report"`
}

func (xmlConverter) Convert(reports ...validation.Report) ([]byte, error) {
	d, err := xml.Marshal(response{reports})
	return d, err
}
