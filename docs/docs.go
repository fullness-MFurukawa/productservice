// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://localhost:8081",
        "contact": {
            "name": "Fullness,Inc.",
            "url": "https://www.fullness.co.jp",
            "email": "furukawa@fullness.co.jp"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.apache.org/licenses/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/add": {
            "post": {
                "description": "新商品を登録する",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "新商品を登録する",
                "parameters": [
                    {
                        "description": "商品名、単価、カテゴリ番号",
                        "name": "newproduct",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.ProductDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "Object"
                        }
                    }
                }
            }
        },
        "/change": {
            "put": {
                "description": "商品を更新する",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "商品を更新する",
                "parameters": [
                    {
                        "description": "商品番号、商品名、単価",
                        "name": "newproduct",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.ProductDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "Object"
                        }
                    }
                }
            }
        },
        "/list": {
            "get": {
                "description": "登録されたすべての商品を取得する",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "商品一覧を取得する",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "Object"
                        }
                    }
                }
            }
        },
        "/remove/{id}": {
            "delete": {
                "description": "削除する",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "商品を削除する",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商品番号",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "Object"
                        }
                    }
                }
            }
        },
        "/search": {
            "get": {
                "description": "キーワードで検索した結果を取得する",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "指定されたキーワードで検索した結果を取得する",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商品キーワード",
                        "name": "keyword",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "Object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/result.FailureDto"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "product.ProductDto": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "result.FailureDto": {
            "type": "object",
            "properties": {
                "function": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/product",
	Schemes:          []string{},
	Title:            "商品アクセスAPIサンプル タイプ-1",
	Description:      "商品および商品カテゴリを管理するAPIサービス",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
