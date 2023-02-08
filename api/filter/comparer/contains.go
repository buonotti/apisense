package comparer

import (
	"reflect"

	"github.com/buonotti/apisense/util"
)

type containsComparer struct{}

func (containsComparer) Compare(a any, b any) bool {
	if reflect.TypeOf(a) != reflect.TypeOf([]any{}) {
		return false
	}

	array, _ := a.([]any)
	if len(array) == 0 {
		return false
	}

	childArray, isArray := array[0].([]any)
	if isArray {
		return util.Any(childArray, func(item any) bool {
			return containsComparer{}.Compare(item, b)
		})
	} else {
		stringArray := util.Map(array, func(item any) string {
			return item.(string)
		})
		return util.Contains(stringArray, b.(string))
	}
}
