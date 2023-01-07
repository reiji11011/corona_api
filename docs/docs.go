// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/patient/details/": {
            "get": {
                "description": "指定した期間と都道府県の感染者数情報を取得する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patients"
                ],
                "summary": "感染者数詳細リスト取得",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 20230101,
                        "description": "開始日",
                        "name": "start_date",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 20230102,
                        "description": "終了日",
                        "name": "end_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "\"北海道\"",
                        "description": "都道府県名",
                        "name": "area",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
