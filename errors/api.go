package errors

import (
	"github.com/joomcode/errorx"
)

var ApiErrors = errorx.NewNamespace("api")
var InvalidWhereClauseError = ApiErrors.NewType("invalid_where_clause")
