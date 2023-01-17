package controllers

import (
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/tidwall/gjson"

	"github.com/buonotti/apisense/conversion"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation"
)

var operators = map[string]comparer{
	".eq.":        eqComparer{},
	".ne.":        neComparer{},
	".gt.":        gtComparer{},
	".gte.":       gteComparer{},
	".lt.":        ltComparer{},
	".lte.":       lteComparer{},
	".contains.":  containsComparer{},
	".ncontains.": ncontainsComparer{},
}

type comparer interface {
	compare(a any, b any) bool
}

type eqComparer struct{}

func (eqComparer) compare(a any, b any) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) && reflect.TypeOf(a) != reflect.TypeOf([]any{}) {
		return false
	}
	switch a.(type) {
	case []any:
		return util.Any(a.([]any), func(item any) bool {
			return eqComparer{}.compare(item, b)
		})
	default:
		return a == b
	}
}

type neComparer struct{}

func (neComparer) compare(a any, b any) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) && reflect.TypeOf(a) != reflect.TypeOf([]any{}) {
		return true
	}

	switch a.(type) {
	case []any:
		return util.All(a.([]any), func(item any) bool {
			return neComparer{}.compare(item, b)
		})
	default:
		return a != b
	}
}

type gtComparer struct{}

func (gtComparer) compare(a any, b any) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) && reflect.TypeOf(a) != reflect.TypeOf([]any{}) {
		return false
	}
	switch a.(type) {
	case string:
		return a.(string) > b.(string)
	case float64:
		return a.(float64) > b.(float64)
	case time.Time:
		return a.(time.Time).After(b.(time.Time)) || a.(time.Time).Equal(b.(time.Time))
	case []any:
		return util.Any(a.([]any), func(item any) bool {
			return gtComparer{}.compare(item, b)
		})
	}
	return false
}

type gteComparer struct{}

func (gteComparer) compare(a any, b any) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) && reflect.TypeOf(a) != reflect.TypeOf([]any{}) {
		return false
	}
	switch a.(type) {
	case string:
		return a.(string) >= b.(string)
	case float64:
		return a.(float64) >= b.(float64)
	case time.Time:
		return a.(time.Time).After(b.(time.Time)) || a.(time.Time).Equal(b.(time.Time))
	case []any:
		return util.Any(a.([]any), func(item any) bool {
			return gteComparer{}.compare(item, b)
		})
	}
	return false
}

type ltComparer struct{}

func (ltComparer) compare(a any, b any) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) && reflect.TypeOf(a) != reflect.TypeOf([]any{}) {
		return false
	}
	switch a.(type) {
	case string:
		return a.(string) < b.(string)
	case float64:
		return a.(float64) < b.(float64)
	case time.Time:
		return a.(time.Time).After(b.(time.Time)) || a.(time.Time).Equal(b.(time.Time))
	case []any:
		return util.Any(a.([]any), func(item any) bool {
			return ltComparer{}.compare(item, b)
		})
	}
	return false
}

type lteComparer struct{}

func (lteComparer) compare(a any, b any) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) && reflect.TypeOf(a) != reflect.TypeOf([]any{}) {
		return false
	}
	switch a.(type) {
	case string:
		return a.(string) <= b.(string)
	case float64:
		return a.(float64) <= b.(float64)
	case time.Time:
		return a.(time.Time).Before(b.(time.Time)) || a.(time.Time).Equal(b.(time.Time))
	case []any:
		return util.Any(a.([]any), func(item any) bool {
			return lteComparer{}.compare(item, b)
		})
	}
	return false
}

type containsComparer struct{}

func (containsComparer) compare(a any, b any) bool {
	if reflect.TypeOf(a) != reflect.TypeOf([]any{}) {
		return false
	}
	arr, _ := a.([]any)
	if len(arr) == 0 {
		return false
	}
	subArr, ok := arr[0].([]any)
	if ok {
		return util.Any(subArr, func(item any) bool {
			return containsComparer{}.compare(item, b)
		})
	} else {
		strArr, ok := arr[0].(string)
		if ok {
			return strings.Contains(strArr, b.(string))
		} else {
			return false
		}
	}
}

type ncontainsComparer struct{}

func (ncontainsComparer) compare(a any, b any) bool {
	if reflect.TypeOf(a) != reflect.TypeOf([]any{}) {
		return false
	}
	return !util.Contains(a.([]string), b.(string))
}

// AllReports godoc
// @Summary Get all the reports
// @Description Gets a list of all reports that can be filtered with a query
// @ID all-reports
// @Tags reports
// @Param where query string false "field.op.value"
// @Param format query string false "json"
// @Success 200
// @Failure 500
// @Router /api/reports [get]
func AllReports(c *gin.Context) {
	allReports, err := validation.Reports()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	allReports, err = filterReports(c, allReports)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	writeFormattedReport(c, allReports...)
}

func filterReports(c *gin.Context, reports []validation.Report) ([]validation.Report, error) {
	whereClause := c.Query("where")
	if whereClause == "" {
		return reports, nil
	}

	whereClause = strings.ReplaceAll(whereClause, "$", "#")

	if !util.Any(util.Keys(operators), func(operator string) bool { return strings.Contains(whereClause, operator) }) {
		return nil, errors.InvalidWhereClauseError.New("invalid where clause. where clause does not contain any valid operator")
	}

	reports, err := validation.Reports()
	if err != nil {
		return nil, err
	}

	filterPredicate, err := buildFilterPredicate(whereClause)
	if err != nil {
		return nil, err
	}
	return util.Where(reports, filterPredicate), nil
}

func buildFilterPredicate(whereClause string) (func(validation.Report) bool, error) {
	op := util.FindFirst(util.Keys(operators), func(op string) bool { return strings.Contains(whereClause, op) })
	if op == nil {
		return nil, errors.InvalidWhereClauseError.New("invalid where clause. where clause does not contain any valid operator")
	}

	comp := operators[*op]

	key := strings.Split(whereClause, *op)[0]
	value := strings.Split(whereClause, *op)[1]
	return func(report validation.Report) bool {
		jsonString, err := json.Marshal(report)
		errors.HandleError(err)
		data := gjson.GetBytes(jsonString, key)
		if strings.Contains(strings.ToLower(key), "time") {
			t, err := time.Parse("2006-01-02T15:04:05.000Z", value)
			if err != nil {
				log.ApiLogger.Error(err.Error())
			}
			return comp.compare(data.Time(), t)
		}
		return comp.compare(data.Value(), value)
	}, nil
}

// @BasePath /api

// Report godoc
// @Summary Get one report
// @Description Gets a single report identified by his id
// @ID report
// @Tags reports
// @Param format query string false "json"
// @Param id path string true "qNg8rJX"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/reports/:id [get]
func Report(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"message": "id is required"})
		return
	}
	reports, err := validation.Reports()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	report := util.FindFirst(reports, func(report validation.Report) bool {
		return report.Id == id
	})
	if report == nil {
		c.AbortWithStatusJSON(404, gin.H{"message": "report not found"})
		return
	}
	writeFormattedReport(c, *report)
}

func writeFormattedReport(c *gin.Context, reports ...validation.Report) {
	body := strings.Builder{}
	format := c.Query("format")
	if format == "" {
		format = "json"
	}
	formatter := conversion.Get(format)
	if formatter == nil {
		c.AbortWithStatusJSON(500, gin.H{"message": "Unknown format: " + format})
		return
	}
	d, err := formatter.Convert(reports...)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
	}
	body.Write(d)
	c.Data(200, "application/"+format, []byte(body.String()))
}
