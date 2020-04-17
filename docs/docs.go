// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "devtraders",
            "url": "https://dev4traders.com",
            "email": "mikhail@dev4traders.com"
        },
        "license": {
            "name": "Commerce"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "get user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Auth user",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserLoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/trades/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "add by json trade",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trades"
                ],
                "summary": "Add an trade",
                "parameters": [
                    {
                        "description": "Add trade",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddTrade"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Trade"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/trades/close": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "close by json trade",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trades"
                ],
                "summary": "Close an trade",
                "parameters": [
                    {
                        "description": "Close trade",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CloseTrade"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/trades/update": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update by json trade",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trades"
                ],
                "summary": "Update an trade",
                "parameters": [
                    {
                        "description": "Update trade",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateTrade"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/trades/{login}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get User trades by login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trades"
                ],
                "summary": "List an user trades",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "login",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Trade"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httputil.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "model.AddTrade": {
            "type": "object",
            "required": [
                "command",
                "price",
                "symbol",
                "volume"
            ],
            "properties": {
                "command": {
                    "type": "integer",
                    "example": 0
                },
                "price": {
                    "type": "number",
                    "example": 1.1456
                },
                "sl": {
                    "type": "number",
                    "example": 1.1456
                },
                "symbol": {
                    "type": "string",
                    "example": "EURUSD"
                },
                "tp": {
                    "type": "number",
                    "example": 1.1456
                },
                "volume": {
                    "type": "number",
                    "example": 0.1
                }
            }
        },
        "model.CloseTrade": {
            "type": "object",
            "required": [
                "ticket"
            ],
            "properties": {
                "ticket": {
                    "type": "integer",
                    "example": 101
                },
                "volume": {
                    "type": "number",
                    "example": 0.1
                }
            }
        },
        "model.LoginUser": {
            "type": "object",
            "required": [
                "login",
                "password"
            ],
            "properties": {
                "login": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "password": {
                    "type": "string",
                    "example": "1234Aa"
                }
            }
        },
        "model.Trade": {
            "type": "object",
            "properties": {
                "close_price": {
                    "type": "number"
                },
                "close_time": {
                    "type": "integer"
                },
                "cmd": {
                    "type": "integer"
                },
                "comment": {
                    "type": "string"
                },
                "digits": {
                    "type": "integer"
                },
                "expiration": {
                    "type": "integer"
                },
                "login": {
                    "type": "integer"
                },
                "magic": {
                    "type": "integer"
                },
                "open_price": {
                    "type": "number"
                },
                "open_time": {
                    "type": "integer"
                },
                "profit": {
                    "type": "number"
                },
                "sl": {
                    "type": "number"
                },
                "symbol": {
                    "type": "string"
                },
                "ticket": {
                    "type": "integer"
                },
                "tp": {
                    "type": "number"
                },
                "volume": {
                    "type": "integer"
                }
            }
        },
        "model.UpdateTrade": {
            "type": "object",
            "required": [
                "ticket"
            ],
            "properties": {
                "price": {
                    "type": "number",
                    "example": 1.1456
                },
                "sl": {
                    "type": "number",
                    "example": 1.1456
                },
                "ticket": {
                    "type": "integer",
                    "example": 101
                },
                "tp": {
                    "type": "number",
                    "example": 1.1456
                }
            }
        },
        "model.UserLoginResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "expire": {
                    "type": "string"
                },
                "token": {
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

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "MT4 Trader REST API",
	Description: "This is a mt4 trader rest api server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
