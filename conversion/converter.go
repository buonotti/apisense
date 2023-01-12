package conversion

import (
	"github.com/buonotti/odh-data-monitor/util"
	"github.com/buonotti/odh-data-monitor/validation"
)

type Converter interface {
	Convert(report ...validation.Report) ([]byte, error)
	//ConvertMany(reports []validation.Report) ([]byte, error)
}

var convMap = map[string]Converter{
	"json": Json(),
	"csv":  Csv(),
	"xml":  Xml(),
}

func Get(name string) Converter {
	return convMap[name]
}

func Converters() []string {
	return util.Keys(convMap)
}
