// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Gets a list of all definitions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "definitions"
                ],
                "summary": "Get all the definitions",
                "operationId": "all-definitions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                    "201": {
                        "description": "Created",
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Gets a single definition identified by his endpoint name",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "definitions"
                ],
                "summary": "Get one definition",
                "operationId": "definition",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                "produces": [
                    "application/json"
                ],
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
        "/login": {
            "post": {
                "description": "Logs a user in using the provided credentials",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Logs a user in",
                "operationId": "login-user",
                "parameters": [
                    {
                        "description": "content",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
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
        "/reports": {
            "get": {
                "description": "Gets a list of all reports that can be filtered with a query",
                "produces": [
                    "application/json"
                ],
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
                "produces": [
                    "application/json"
                ],
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
        }
    },
    "definitions": {
        "controllers.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "Holds the information about what happened",
                    "type": "string"
                }
            }
        },
        "controllers.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
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
                "authorization": {
                    "description": "Authorization is the value to set for the authorization header",
                    "type": "string"
                },
                "baseUrl": {
                    "description": "BaseUrl is the base path of the endpoint",
                    "type": "string"
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
                "headers": {
                    "description": "Headers are additional headers to set for the request",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "jwt_login": {
                    "description": "JwtLogin are options to auto-get a login token for a request.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/definitions.JwtLoginOptions"
                        }
                    ]
                },
                "method": {
                    "description": "Method is the name of the http-method to use for the request",
                    "type": "string"
                },
                "name": {
                    "description": "Name is the name of the endpoint",
                    "type": "string"
                },
                "ok_code": {
                    "description": "The expected status code",
                    "type": "integer"
                },
                "payload": {
                    "description": "Payload is the payload to use in case of a POST or PUT request",
                    "type": "object",
                    "additionalProperties": {}
                },
                "queryParameters": {
                    "description": "QueryParameters are all the query parameters that should be added to the call",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/definitions.QueryDefinition"
                    }
                },
                "responseSchema": {
                    "description": "ResponseSchema describes how the response should look like"
                },
                "test_case_names": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "variables": {
                    "description": "Variables are all the variables that should be interpolated in the base url and the query parameters",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/definitions.Variable"
                    }
                },
                "version": {
                    "description": "Version is the version of the definition",
                    "type": "integer"
                }
            }
        },
        "definitions.JwtLoginOptions": {
            "type": "object",
            "required": [
                "url"
            ],
            "properties": {
                "login_payload": {
                    "description": "LoginPayload is the json or yml payload to send",
                    "type": "object",
                    "additionalProperties": {}
                },
                "token_key_name": {
                    "description": "TokenKeyName is the name of the key in the response which contains the token",
                    "type": "string"
                },
                "url": {
                    "description": "Url is the url to the login endpoint",
                    "type": "string"
                }
            }
        },
        "definitions.QueryDefinition": {
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
                "name": {
                    "description": "Name is the name of the test case result",
                    "type": "string"
                },
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
                    "allOf": [
                        {
                            "$ref": "#/definitions/validators.ValidatorStatus"
                        }
                    ]
                }
            }
        },
        "validators.ValidatorStatus": {
            "type": "string",
            "enum": [
                "unknown",
                "success",
                "skipped",
                "fail"
            ],
            "x-enum-varnames": [
                "ValidatorStatusUnknown",
                "ValidatorStatusSuccess",
                "ValidatorStatusSkipped",
                "ValidatorStatusFail"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Apisense API",
	Description:      "Api specification for the Apisense API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
