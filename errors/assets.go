package errors

import (
	"github.com/joomcode/errorx"
)

var AssetErrors = errorx.NewNamespace("asset")
var CannotLoadAssetError = AssetErrors.NewType("cannot_load_asset")
