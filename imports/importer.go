package imports

import "github.com/buonotti/apisense/validation/definitions"

// Importer is something capable of converting a given file and its content to a list of definitions
type Importer interface {
	Import(file string, content []byte) ([]definitions.Endpoint, error)
}
