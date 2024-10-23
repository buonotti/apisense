package filter

import (
	"strings"
	"time"

	"github.com/buonotti/apisense/api/filter/comparer"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/util"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

// Parse parses the given query and creates an appropriate filter function
func Parse[T any](query string) (func(T) bool, error) {
	comparerType := comparer.ExtractOperator(query)
	if comparerType == "" {
		return nil, errors.InvalidWhereClauseError.New("invalid where clause (no valid operator found): %s", query)
	}

	comp := comparer.New(comparerType)

	key := strings.Split(query, comparerType)[0]
	value := strings.Split(query, comparerType)[1]
	filterPredicate := func(item T) bool {
		jsonString, err := json.Marshal(item)
		if err != nil {
			log.ApiLogger().Fatal(err)
		}
		data := gjson.GetBytes(jsonString, key)
		if strings.Contains(strings.ToLower(key), "time") {
			parsedTime, err := time.Parse(util.ApisenseTimeFormat, value)
			if err != nil {
				log.ApiLogger().Error(err.Error())
			}
			return comp.Compare(data.Time(), parsedTime)
		}
		return comp.Compare(data.Value(), value)
	}

	return filterPredicate, nil
}

// ParseFromContext is the same as Parse it just auto-extracts the query from the given fiber.Ctx
func ParseFromContext[T any](c *fiber.Ctx) (func(T) bool, error) {
	whereClause := c.Query("where")
	if whereClause == "" {
		return func(_ T) bool { return true }, nil
	}

	whereClause = strings.ReplaceAll(whereClause, "$", "#")
	return Parse[T](whereClause)
}
