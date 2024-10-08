package errors

import (
	"github.com/joomcode/errorx"
)

var (
	ApiErrors                    = errorx.NewNamespace("api")
	InvalidWhereClauseError      = ApiErrors.NewType("invalid_where_clause")
	CannotStopApiServiceError    = ApiErrors.NewType("cannot_stop_api_service")
	CannotUpgradeWebsocketError  = ApiErrors.NewType("cannot_upgrade_websocket", fatalTrait)
	CannotWriteWebsocketError    = ApiErrors.NewType("cannot_write_websocket")
	CannotFindReportError        = ApiErrors.NewType("cannot_find_report")
	IdRequiredError              = ApiErrors.NewType("id_required")
	NameRequiredError            = ApiErrors.NewType("name_required")
	CannotFindDefinitionError    = ApiErrors.NewType("cannot_find_definition")
	DefinitionAlreadyExistsError = ApiErrors.NewType("definition_already_exists")
	TokenError                   = ApiErrors.NewType("token_error")
	LoginError                   = ApiErrors.NewType("login_error")
	UserAlreadyExistsError       = ApiErrors.NewType("user_already_exists", fatalTrait)
	MissingSigningKeyError       = ApiErrors.NewType("signing_key_missing", fatalTrait)
)
