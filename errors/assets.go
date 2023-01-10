package errors

import (
	"github.com/joomcode/errorx"
)

// AssetErrors is the namespace holding all asset loading related error
var AssetErrors = errorx.NewNamespace("asset")

// CannotLoadAssetError is returned when asset loading fails
var CannotLoadAssetError = AssetErrors.NewType("cannot_load_asset")
