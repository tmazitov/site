// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Linggar Primahastoko",
            "url": "https://linggar.asia",
            "email": "x@linggar.asia"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/entry": {
            "post": {
                "description": "User auth",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User auth",
                "operationId": "sign-in",
                "parameters": [
                    {
                        "description": "SignIn params",
                        "name": "Params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.SignInParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "User not exists",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/list": {
            "get": {
                "description": "User list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User list",
                "operationId": "list",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 16,
                        "name": "per_page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "name": "timestamp",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/new": {
            "post": {
                "description": "Create new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create new user",
                "operationId": "sign-up",
                "parameters": [
                    {
                        "description": "SignUp params",
                        "name": "Params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.SignUpParams"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "This username/email is taken",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/profile": {
            "get": {
                "description": "User profile data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User profile data",
                "operationId": "profile",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/refresh": {
            "put": {
                "description": "Refresh user token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Refresh user token",
                "operationId": "refresh",
                "responses": {
                    "200": {
                        "description": "Success created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/upgrade": {
            "get": {
                "description": "Upgrade user role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Upgrade user role",
                "operationId": "upgrade",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "user.SignInParams": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "123456789"
                },
                "username": {
                    "type": "string",
                    "example": "example"
                }
            }
        },
        "user.SignUpParams": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "example@gmail.com"
                },
                "password": {
                    "type": "string",
                    "minLength": 8,
                    "example": "123456789"
                },
                "username": {
                    "type": "string",
                    "minLength": 5,
                    "example": "example"
                }
            }
        }
    },
    "securityDefinitions": {
        "": {
            "type": "basic"
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
	Host:        "localhost:8000/api",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Go Restful API with Swagger",
	Description: "Simple swagger implementation in Go HTTP",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
