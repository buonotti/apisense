package filter

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"

	"github.com/buonotti/apisense/api/filter/comparer"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/util"
)

type Filter[T any] struct {
	predicate func(T) bool
}

func New[T any]() *Filter[T] {
	return &Filter[T]{
		predicate: func(T) bool {
			return true
		},
	}
}

func Parse[T any](query string) (*Filter[T], error) {
	comparerType := comparer.ExtractOperator(query)
	if comparerType == "" {
		return nil, errors.InvalidWhereClauseError.New("invalid where clause (no valid operator found): %s", query)
	}

	comp := comparer.New(comparerType)

	key := strings.Split(query, comparerType)[0]
	value := strings.Split(query, comparerType)[1]
	filterPredicate := func(item T) bool {
		jsonString, err := json.Marshal(item)
		errors.HandleError(err)
		data := gjson.GetBytes(jsonString, key)
		if strings.Contains(strings.ToLower(key), "time") {
			parsedTime, err := time.Parse("2006-01-02T15:04:05.000Z", value)
			if err != nil {
				log.ApiLogger.Error(err.Error())
			}
			return comp.Compare(data.Time(), parsedTime)
		}
		return comp.Compare(data.Value(), value)
	}

	return &Filter[T]{
		predicate: filterPredicate,
	}, nil
}

func ParseFromContext[T any](c *gin.Context) (*Filter[T], error) {
	whereClause := c.Query("where")
	if whereClause == "" {
		return &Filter[T]{
			predicate: func(T) bool {
				return true
			},
		}, nil
	}

	whereClause = strings.ReplaceAll(whereClause, "$", "#")
	return Parse[T](whereClause)
}

func (f *Filter[T]) Apply(items []T) []T {
	return util.Where(items, f.predicate)
}
