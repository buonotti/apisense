package comparer

import (
	"time"

	"github.com/buonotti/apisense/v2/util"
	"github.com/goccy/go-reflect"
)

type gtComparer struct{}

func (gtComparer) Compare(a any, b any) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) && reflect.TypeOf(a) != reflect.TypeOf([]any{}) {
		return false
	}

	switch a.(type) {
	case string:
		return a.(string) > b.(string)
	case float64:
		return a.(float64) > b.(float64)
	case time.Time:
		aTime := a.(time.Time)
		bTime := b.(time.Time)
		return aTime.After(bTime)
	case []any:
		return util.Any(a.([]any), func(item any) bool {
			return gtComparer{}.Compare(item, b)
		})
	}

	return false
}
