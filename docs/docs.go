// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "buonotti",
            "url": "https://github.com/buonotti/apisense/issues"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/definitions": {
            "get": {
                "description": "Gets a list of all definitions",
                "tags": [
                    "definitions"
                ],
                "summary": "Get all the definitions",
                "operationId": "all-definitions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/definitions.Endpoint"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new definition",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "definitions"
                ],
                "summary": "Create a definition",
                "operationId": "create-definition",
                "parameters": [
                    {
                        "description": "Endpoint definition",
                        "name": "definition",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/definitions.Endpoint"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/definitions.Endpoint"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/definitions/:name": {
            "get": {
                "description": "Gets a single definition identified by his endpoint name",
                "tags": [
                    "definitions"
                ],
                "summary": "Get one definition",
                "operationId": "definition",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bluetooth",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/definitions.Endpoint"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Get the health status of the API",
                "tags": [
                    "health"
                ],
                "summary": "Health check",
                "operationId": "health",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/reports": {
            "get": {
                "description": "Gets a list of all reports that can be filtered with a query",
                "tags": [
                    "reports"
                ],
                "summary": "Get all the reports",
                "operationId": "all-reports",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Query in the format: field.op.value (optional)",
                        "name": "where",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Return format: json or csv (default: json)",
                        "name": "format",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/pipeline.Report"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/reports/:id": {
            "get": {
                "description": "Gets a single report identified by his id",
                "tags": [
                    "reports"
                ],
                "summary": "Get one report",
                "operationId": "report",
                "parameters": [
                    {
                        "type": "string",
                        "description": "json",
                        "name": "format",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "qNg8rJX",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pipeline.Report"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/ws": {
            "get": {
                "description": "Connect to this endpoint with the ws:// protocol to instantiate a websocket connection to get updates for new reports",
                "tags": [
                    "reports"
                ],
                "summary": "Open a websocket connection to receive notifications",
                "operationId": "ws",
                "responses": {
                    "101": {
                        "description": "Switching Protocols"
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "definitions.Endpoint": {
            "type": "object",
            "required": [
                "baseUrl",
                "format",
                "name",
                "responseSchema"
            ],
            "properties": {
                "baseUrl": {
                    "description": "BaseUrl is the base path of the endpoint",
                    "type": "string"
                },
                "enabled": {
                    "description": "IsEnabled is a boolean that indicates if the endpoint is enabled (not contained in the definition)",
                    "type": "boolean"
                },
                "excludedValidators": {
                    "description": "ExcludedValidators is a list of validators that should not be used for this endpoint",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "format": {
                    "description": "Format is the response format of the",
                    "type": "string"
                },
                "name": {
                    "description": "Name is the name of the endpoint",
                    "type": "string"
                },
                "queryParameters": {
                    "description": "QueryParameters are all the query parameters that should be added to the call",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/query.Definition"
                    }
                },
                "responseSchema": {
                    "description": "ResponseSchema describes how the response should look like",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/definitions.SchemaEntry"
                    }
                },
                "variables": {
                    "description": "Variables are all the variables that should be interpolated in the base url and the query parameters",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/definitions.Variable"
                    }
                }
            }
        },
        "definitions.SchemaEntry": {
            "type": "object",
            "required": [
                "fields",
                "name",
                "required",
                "type"
            ],
            "properties": {
                "fields": {
                    "description": "Fields describe the children of this field if the field is an object or array",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/definitions.SchemaEntry"
                    }
                },
                "max": {
                    "description": "Maximum is the maximum allowed value of the field"
                },
                "min": {
                    "description": "Minimum is the minimum allowed value of the field"
                },
                "name": {
                    "description": "Name is the name of the field",
                    "type": "string"
                },
                "required": {
                    "description": "Required is true if the field is required (not null or not empty in case of an array)",
                    "type": "boolean"
                },
                "type": {
                    "description": "Type is the type of the field",
                    "type": "string"
                }
            }
        },
        "definitions.Variable": {
            "type": "object",
            "required": [
                "constant",
                "name",
                "values"
            ],
            "properties": {
                "constant": {
                    "description": "IsConstant is true if the value of the variable is constant or else false",
                    "type": "boolean"
                },
                "name": {
                    "description": "Name is the name of the variable",
                    "type": "string"
                },
                "values": {
                    "description": "Values are all the possible values of the variable (only 1 in case of a constant)",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "pipeline.Report": {
            "type": "object",
            "properties": {
                "endpoints": {
                    "description": "Endpoints is a collection of ValidatedEndpoint holding the validation results",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pipeline.ValidatedEndpoint"
                    }
                },
                "id": {
                    "description": "Id is a unique identifier for each report",
                    "type": "string"
                },
                "time": {
                    "description": "Time is the timestamp of the report",
                    "type": "string"
                }
            }
        },
        "pipeline.TestCaseResult": {
            "type": "object",
            "properties": {
                "url": {
                    "description": "Url is the url of the api call (with query parameters)",
                    "type": "string"
                },
                "validatorResults": {
                    "description": "ValidatorResults is the collection of ValidatorResult that describe the result of each validator",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pipeline.ValidatorResult"
                    }
                }
            }
        },
        "pipeline.ValidatedEndpoint": {
            "type": "object",
            "properties": {
                "endpointName": {
                    "description": "EndpointName is he name of the endpoint",
                    "type": "string"
                },
                "testCaseResults": {
                    "description": "TestCaseResults are the collection of TestCaseResult that describe the result of validating a single api call",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pipeline.TestCaseResult"
                    }
                }
            }
        },
        "pipeline.ValidatorResult": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "Message is the error message of the validator",
                    "type": "string"
                },
                "name": {
                    "description": "Name is the name of the validator",
                    "type": "string"
                },
                "status": {
                    "description": "Status is the status of the validator (success/fail/skipped)",
                    "type": "string"
                }
            }
        },
        "query.Definition": {
            "type": "object",
            "required": [
                "name",
                "value"
            ],
            "properties": {
                "name": {
                    "description": "Name is the name of the query parameter",
                    "type": "string"
                },
                "value": {
                    "description": "Value is the value of the query parameter",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Apisense API",
	Description:      "Api specification for the Apisense API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
