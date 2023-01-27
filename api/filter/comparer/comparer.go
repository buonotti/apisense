package comparer

import (
	"strings"
)

var operators = map[string]Comparer{
	".eq.":       eqComparer{},
	".ne.":       neComparer{},
	".gt.":       gtComparer{},
	".gte.":      gteComparer{},
	".lt.":       ltComparer{},
	".lte.":      lteComparer{},
	".contains.": containsComparer{},
	".excludes.": excludesComparer{},
}

// Comparer is an interface to compare any two values
type Comparer interface {
	Compare(a any, b any) bool
}

// New returns a new comparer of the given type t
func New(operatorType string) Comparer {
	return operators[operatorType]
}

// ExtractOperator extracts the operator type from a given string
func ExtractOperator(filterQuery string) string {
	for operatorType := range operators {
		if strings.Contains(filterQuery, operatorType) {
			return operatorType
		}
	}

	return ""
}
