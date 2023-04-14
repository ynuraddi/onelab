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
        "/login": {
            "post": {
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User login input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginUserRq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "Create a new user with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User information",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserCreateRq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "user created",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    },
                    "422": {
                        "description": "user already exist",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get user by id in query param",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    },
                    "401": {
                        "description": "missing or malformed jwt",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    },
                    "404": {
                        "description": "user is not exist",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete user by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user deleted",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    },
                    "401": {
                        "description": "missing or malformed jwt",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    },
                    "404": {
                        "description": "user is not exist",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update user by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User information",
                        "name": "input",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.UserUpdateRq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user updated",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    },
                    "401": {
                        "description": "missing or malformed jwt",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    },
                    "404": {
                        "description": "user is not exist",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.Envelope"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.Envelope": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "model.LoginUserRq": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "minLength": 5
                },
                "password": {
                    "type": "string",
                    "minLength": 5
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.UserCreateRq": {
            "type": "object",
            "required": [
                "login",
                "name",
                "password"
            ],
            "properties": {
                "login": {
                    "type": "string",
                    "minLength": 5
                },
                "name": {
                    "type": "string",
                    "minLength": 5
                },
                "password": {
                    "type": "string",
                    "minLength": 5
                }
            }
        },
        "model.UserUpdateRq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer",
                    "minimum": 1
                },
                "login": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Onelab API",
	Description:      "Onelab project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
