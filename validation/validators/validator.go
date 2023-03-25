package validators

import "github.com/buonotti/apisense/validation/fetcher"

type Validator interface {
	Name() string
	Validate(item fetcher.TestCase) error
	IsFatal() bool
}