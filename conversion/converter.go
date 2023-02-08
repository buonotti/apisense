package conversion

import (
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation"
)

type Converter interface {
	Convert(report ...validation.Report) ([]byte, error)
}

var convMap = map[string]Converter{
	"json": Json(),
	"csv":  Csv(),
}

func Get(name string) Converter {
	return convMap[name]
}

func Converters() []string {
	return util.Keys(convMap)
}
