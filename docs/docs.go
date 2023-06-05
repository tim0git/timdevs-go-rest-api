// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "returns the status of the service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "health check endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.HealthStatus"
                        }
                    }
                }
            }
        },
        "/vehicle": {
            "post": {
                "description": "register a new vehicle",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vehicle"
                ],
                "summary": "register a new vehicle",
                "parameters": [
                    {
                        "description": "Vehicle information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vehicle.Vehicle"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/vehicle/{vin}": {
            "get": {
                "description": "retrieve a vehicle",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vehicle"
                ],
                "summary": "retrieve a vehicle",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Vehicle identification number",
                        "name": "vin",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vehicle.Vehicle"
                        }
                    }
                }
            },
            "patch": {
                "description": "update an existing vehicle",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vehicle"
                ],
                "summary": "update an existing vehicle",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Vehicle identification number",
                        "name": "vin",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Vehicle information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vehicle.Update"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.HealthStatus": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "vehicle.Capacity": {
            "type": "object",
            "required": [
                "unit",
                "value"
            ],
            "properties": {
                "unit": {
                    "type": "string"
                },
                "value": {
                    "type": "integer"
                }
            }
        },
        "vehicle.Update": {
            "type": "object",
            "required": [
                "capacity",
                "color",
                "license_plate",
                "manufacturer",
                "model",
                "year"
            ],
            "properties": {
                "capacity": {
                    "$ref": "#/definitions/vehicle.Capacity"
                },
                "color": {
                    "type": "string"
                },
                "license_plate": {
                    "type": "string"
                },
                "manufacturer": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "vehicle.Vehicle": {
            "type": "object",
            "required": [
                "capacity",
                "color",
                "manufacturer",
                "model",
                "vin",
                "year"
            ],
            "properties": {
                "capacity": {
                    "$ref": "#/definitions/vehicle.Capacity"
                },
                "color": {
                    "type": "string"
                },
                "license_plate": {
                    "type": "string"
                },
                "manufacturer": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "vin": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
