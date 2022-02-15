// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Mustafa Erdem Köşk",
            "email": "erdemkosk@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/config/": {
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Create config with using key type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "config"
                ],
                "summary": "Create config with using key and type",
                "parameters": [
                    {
                        "description": "Config Value",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ConfigInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ExampleSuccessConfig"
                        }
                    }
                }
            }
        },
        "/api/config/{key}": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Get config with using key",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "config"
                ],
                "summary": "Get config with using key",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Config key value",
                        "name": "key",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "Array",
                            "Boolean",
                            "Object",
                            "String",
                            "Number"
                        ],
                        "type": "string",
                        "description": "type",
                        "name": "types",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ExampleGetConfig"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Update config with using key",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "config"
                ],
                "summary": "Update config with using key",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Config key value",
                        "name": "key",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Config Value",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ConfigUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ExampleSuccessConfig"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Delete config with using key",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "config"
                ],
                "summary": "Delete config with using key",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Config key value",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ExampleSuccessConfig"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Get config with using key",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Post config with using key",
                "parameters": [
                    {
                        "description": "Cridentinials",
                        "name": "cridentinials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AuthToken"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AuthToken": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.ConfigFormatted": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "value": {}
            }
        },
        "models.ConfigInput": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "value": {}
            }
        },
        "models.ConfigUpdate": {
            "type": "object",
            "properties": {
                "type": {
                    "type": "string"
                },
                "value": {}
            }
        },
        "models.ExampleGetConfig": {
            "type": "object",
            "properties": {
                "config": {
                    "$ref": "#/definitions/models.ConfigFormatted"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.ExampleSuccessConfig": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.LoginInput": {
            "type": "object",
            "properties": {
                "identity": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Authorization": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Config Service",
	Description:      "Swagger structure prepared for config service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
