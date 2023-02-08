package comparer

import (
	"reflect"
)

type eqComparer struct{}

func (eqComparer) Compare(a any, b any) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}

	return a == b
}
