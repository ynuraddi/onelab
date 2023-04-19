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
        "/book": {
            "post": {
                "description": "Create a new book with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book"
                ],
                "summary": "Create a new book",
                "parameters": [
                    {
                        "description": "Book information",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateBookRq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Book created",
                        "schema": {
                            "$ref": "#/definitions/handler.MsgEnvelope"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "422": {
                        "description": "book already exist",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    }
                }
            }
        },
        "/book/borrow": {
            "post": {
                "description": "Create a new book borrow with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book_borrow"
                ],
                "summary": "Create a new book borrow",
                "parameters": [
                    {
                        "description": "Book borrow information time:2020-04-17T18:25:43.511Z",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateBookBorrowRq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "record created",
                        "schema": {
                            "$ref": "#/definitions/handler.MsgEnvelope"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    }
                }
            }
        },
        "/book/borrow/debtor/list": {
            "get": {
                "description": "List borrow debtor record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book_borrow"
                ],
                "summary": "List Borrow debtors",
                "responses": {
                    "302": {
                        "description": "Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.BookBorrowDebtorRp"
                            }
                        }
                    },
                    "401": {
                        "description": "missing or malformed jwt",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "404": {
                        "description": "record is not exist",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    }
                }
            }
        },
        "/book/borrow/metric/list/{id}": {
            "get": {
                "description": "List borrow debtor record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book_borrow"
                ],
                "summary": "List Borrow metric",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Month",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.BookBorrowMetricRp"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "401": {
                        "description": "missing or malformed jwt",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "404": {
                        "description": "record is not exist",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    }
                }
            }
        },
        "/book/borrow/{id}": {
            "get": {
                "description": "Get borrow record by id in query param",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book_borrow"
                ],
                "summary": "Get Borrow record by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "BookBorrow ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/model.BookBorrow"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "401": {
                        "description": "missing or malformed jwt",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "404": {
                        "description": "record is not exist",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book_borrow"
                ],
                "summary": "Delete book borrow record by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "BookBorrow ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "record deleted",
                        "schema": {
                            "$ref": "#/definitions/handler.MsgEnvelope"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "401": {
                        "description": "missing or malformed jwt",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "404": {
                        "description": "record is not exist",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book_borrow"
                ],
                "summary": "Update book borrow by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "BookBorrow ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "record update information",
                        "name": "input",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.UpdateBookBorrowRq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "record updated",
                        "schema": {
                            "$ref": "#/definitions/handler.MsgEnvelope"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "401": {
                        "description": "missing or malformed jwt",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "404": {
                        "description": "record is not exist",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "409": {
                        "description": "edit conflict",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    }
                }
            }
        },
        "/book/{id}": {
            "get": {
                "description": "Get book by id in query param",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book"
                ],
                "summary": "Get book by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "401": {
                        "description": "missing or malformed jwt",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "404": {
                        "description": "book is not exist",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book"
                ],
                "summary": "Delete book by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "book deleted",
                        "schema": {
                            "$ref": "#/definitions/handler.MsgEnvelope"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "401": {
                        "description": "missing or malformed jwt",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "404": {
                        "description": "book is not exist",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book"
                ],
                "summary": "Update book by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Book information",
                        "name": "input",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.UpdateBookRq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "book updated",
                        "schema": {
                            "$ref": "#/definitions/handler.MsgEnvelope"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "401": {
                        "description": "missing or malformed jwt",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "404": {
                        "description": "book is not exist",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "409": {
                        "description": "edit conflict",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    }
                }
            }
        },
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
                            "$ref": "#/definitions/model.LogInRq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "logined",
                        "schema": {
                            "$ref": "#/definitions/handler.MsgEnvelope"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
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
                            "$ref": "#/definitions/model.CreateUserRq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "user created",
                        "schema": {
                            "$ref": "#/definitions/handler.MsgEnvelope"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "422": {
                        "description": "user already exist",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
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
                        "description": "User ID",
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
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "401": {
                        "description": "missing or malformed jwt",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "404": {
                        "description": "user is not exist",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
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
                            "$ref": "#/definitions/handler.MsgEnvelope"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "401": {
                        "description": "missing or malformed jwt",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "404": {
                        "description": "user is not exist",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
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
                            "$ref": "#/definitions/model.UpdateUserRq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user updated",
                        "schema": {
                            "$ref": "#/definitions/handler.MsgEnvelope"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "401": {
                        "description": "missing or malformed jwt",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "404": {
                        "description": "user is not exist",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "409": {
                        "description": "edit conflict",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrEnvelope"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.ErrEnvelope": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "handler.MsgEnvelope": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "model.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "model.BookBorrow": {
            "type": "object",
            "properties": {
                "book_id": {
                    "type": "integer"
                },
                "borrow_date": {
                    "type": "string"
                },
                "borrow_id": {
                    "type": "integer"
                },
                "return_date": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "uuid": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "model.BookBorrowDebtorRp": {
            "type": "object",
            "properties": {
                "book_id": {
                    "type": "integer"
                },
                "book_name": {
                    "type": "string"
                },
                "borrow_date": {
                    "type": "string"
                },
                "borrow_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "model.BookBorrowMetricRp": {
            "type": "object",
            "properties": {
                "book_amount": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "model.CreateBookBorrowRq": {
            "type": "object",
            "required": [
                "book_id",
                "borrow_date",
                "user_id"
            ],
            "properties": {
                "book_id": {
                    "type": "integer",
                    "minimum": 1
                },
                "borrow_date": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "model.CreateBookRq": {
            "type": "object",
            "required": [
                "author",
                "price",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string",
                    "minLength": 5
                },
                "price": {
                    "type": "integer",
                    "minimum": 100
                },
                "title": {
                    "type": "string",
                    "minLength": 5
                }
            }
        },
        "model.CreateUserRq": {
            "type": "object",
            "required": [
                "login",
                "password",
                "user_name"
            ],
            "properties": {
                "login": {
                    "type": "string",
                    "minLength": 5
                },
                "password": {
                    "type": "string",
                    "minLength": 5
                },
                "user_name": {
                    "type": "string",
                    "minLength": 5
                }
            }
        },
        "model.LogInRq": {
            "type": "object",
            "required": [
                "login",
                "password"
            ],
            "properties": {
                "login": {
                    "type": "string",
                    "minLength": 5
                },
                "password": {
                    "type": "string",
                    "minLength": 5
                }
            }
        },
        "model.UpdateBookBorrowRq": {
            "type": "object",
            "properties": {
                "book_id": {
                    "type": "integer"
                },
                "borrow_date": {
                    "type": "string"
                },
                "return_date": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.UpdateBookRq": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.UpdateUserRq": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "login": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
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
