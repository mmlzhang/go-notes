package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "An example of gin",
        "title": "Golang Gin API",
        "termsOfService": "https://github.com/cool8sniper/gin-rest-gorm-rbac-sample",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://github.com/cool8sniper/gin-rest-gorm-rbac-sample/blob/master/LICENSE"
        },
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/user/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "IDd",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"id\": 112,\"name\":\"xxx\", \"email\": \"xx@xx.com\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"message\": \"No found user\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}