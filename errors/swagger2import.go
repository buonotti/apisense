package errors

import "github.com/joomcode/errorx"

var (
	SwaggerImportErrors         = errorx.NewNamespace("swaggerimport")
	InvalidRefTypeError         = SwaggerImportErrors.NewType("invalid_ref_type")
	CannotResolveRefError       = SwaggerImportErrors.NewType("cannot_resolve_ref")
	InvalidContentTypeError     = SwaggerImportErrors.NewType("invalid_content_type")
	InvalidSwaggerDocumentError = SwaggerImportErrors.NewType("invalid_document")
	CannotConvertOpenapiV2Spec  = SwaggerImportErrors.NewType("cannot_convert_openapiv2spec")
)
