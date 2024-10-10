package imports

import "github.com/buonotti/apisense/validation/definitions"

type Importer interface {
	Import(file string, content []byte) (definitions.Endpoint, error)
}
