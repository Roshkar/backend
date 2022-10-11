package pkg

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Alexandar Naydenov",
            "email": "alexandar.naydenov99@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/delete/order/{orderId}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Delete an order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the order",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Request has wrong format",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/delete/product/{productId}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Delete a product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the product",
                        "name": "productId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Request has wrong format",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/order": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Get all orders from the shop",
                "responses": {
                    "200": {
                        "description": "Successful request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Submit a new order",
                "parameters": [
                    {
                        "description": "New order details",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.ExampleOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Request has wrong format or not enought quantity of a product",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/order/{orderId}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Get a order by id from the shop",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the order",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Order with such Id not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Update an order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the order",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated order details",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.ExampleOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Request has wrong format",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/product": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get all products from the shop",
                "responses": {
                    "200": {
                        "description": "Successful request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Add a new product",
                "parameters": [
                    {
                        "description": "New product details",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.ExampleProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Request has wrong format",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/product/{productId}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get a product by id from the shop",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the product",
                        "name": "productId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Product with such Id not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Update a product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the product",
                        "name": "productId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated product details",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.ExampleProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Request has wrong format",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "structs.ExampleOrderRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "default": "Sofia Mladost 2"
                },
                "name": {
                    "type": "string",
                    "default": "Ivan Ivanov"
                },
                "phone": {
                    "type": "string",
                    "default": "0888888888"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "id": {
                                "type": "string",
                                "default": "bc264186-9c2e-4533-6ba5-705c160303c1"
                            },
                            "quantity": {
                                "type": "integer",
                                "default": 2
                            }
                        }
                    }
                }
            }
        },
        "structs.ExampleProductRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string",
                    "default": "Men Shirts"
                },
                "name": {
                    "type": "string",
                    "default": "Men Red Shirt"
                },
                "price": {
                    "type": "number",
                    "default": 19.99
                },
                "quantity": {
                    "type": "integer",
                    "default": 1000
                }
            }
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
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Golang Rest Shop Backend",
	Description: "This is a sample rest backend for an online shop.",
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
