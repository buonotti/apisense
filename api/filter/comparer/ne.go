package comparer

import (
	"github.com/goccy/go-reflect"
)

type neComparer struct{}

func (neComparer) Compare(a any, b any) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return true
	}

	return a != b
}
