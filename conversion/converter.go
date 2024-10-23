package conversion

import (
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation/pipeline"
)

// Converter represents a type capable of converting a list of reports to a certain format and then to an []byte
type Converter interface {
	Convert(report ...pipeline.Report) ([]byte, error)
}

var convMap = map[string]Converter{
	"json": Json(),
	"csv":  Csv(),
}

// Get returns a registered converter for the given format
func Get(name string) Converter {
	return convMap[name]
}

// Converters returns all available formats
func Converters() []string {
	return util.Keys(convMap)
}
