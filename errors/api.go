package errors

import (
	"github.com/joomcode/errorx"
)

var ApiErrors = errorx.NewNamespace("api")
var InvalidWhereClauseError = ApiErrors.NewType("invalid_where_clause")
var CannotStopApiServiceError = ApiErrors.NewType("cannot_stop_api_service")
var CannotUpgradeWebsocketError = ApiErrors.NewType("cannot_upgrade_websocket", fatalTrait)
var CannotWriteWebsocketError = ApiErrors.NewType("cannot_write_websocket")
var CannotFindReportError = ApiErrors.NewType("cannot_find_report")
var IdRequiredError = ApiErrors.NewType("id_required")
