package imports

import "github.com/buonotti/apisense/v2/validation/definitions"

// Importer is something capable of converting a given file and its content to a list of definitions
type Importer interface {
	Import(content []byte) ([]definitions.Endpoint, error)
}
