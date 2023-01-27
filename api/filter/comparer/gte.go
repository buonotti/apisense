package comparer

import (
	"reflect"
	"time"

	"github.com/buonotti/apisense/util"
)

type gteComparer struct{}

func (gteComparer) Compare(a any, b any) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) && reflect.TypeOf(a) != reflect.TypeOf([]any{}) {
		return false
	}

	switch a.(type) {
	case string:
		return a.(string) >= b.(string)
	case float64:
		return a.(float64) >= b.(float64)
	case time.Time:
		aTime := a.(time.Time)
		bTime := b.(time.Time)
		return aTime.After(bTime) || aTime.Equal(bTime)
	case []any:
		return util.Any(a.([]any), func(item any) bool {
			return gteComparer{}.Compare(item, b)
		})
	}

	return false
}
