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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/creative-space": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Создание креативной площадки",
                "parameters": [
                    {
                        "description": "Параметры запроса",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.createCreativeSpaceRequestData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ResponseSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/responses.createCreativeSpaceResponseData"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "403": {
                        "description": "Коды ошибок: [1100]",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/v1/health": {
            "get": {
                "summary": "Проверка работоспособности сервера",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseSuccess"
                        }
                    }
                }
            }
        },
        "/api/v1/metro-stations": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Возвращает полный список станций метро",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ResponseSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/responses.getMetroStationsResponseData"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "403": {
                        "description": "Коды ошибок: [1100]",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/v1/session": {
            "get": {
                "summary": "Получение сессии",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseSuccess"
                        }
                    },
                    "403": {
                        "description": "Коды ошибок: [1100]",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
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
                "summary": "Создание сессии пользователя",
                "parameters": [
                    {
                        "description": "Параметры запроса",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.createSessionRequestData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ResponseSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/responses.createSessionResponseData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "summary": "Удаление сессии",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseSuccess"
                        }
                    },
                    "403": {
                        "description": "Коды ошибок: [1100]",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/v1/user": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Создание пользователя",
                "parameters": [
                    {
                        "description": "Параметры запроса",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.createUserRequestData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ResponseSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/responses.createUserResponseData"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "403": {
                        "description": "Коды ошибок: [1100]",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/v1/user/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Возвращает информацию о пользователе",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ResponseSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/responses.getUserResponseData"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "403": {
                        "description": "Коды ошибок: [1100]",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
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
                "summary": "Удаление информации о пользователе",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseSuccess"
                        }
                    },
                    "403": {
                        "description": "Коды ошибок: [1100]",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
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
                "summary": "Обновление информации о пользователе",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Параметры запроса",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.patchUserRequestData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ResponseSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/responses.patchUserResponseData"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "403": {
                        "description": "Коды ошибок: [1100]",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Возвращает список пользователей",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ResponseSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/responses.getUsersResponseData"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "403": {
                        "description": "Коды ошибок: [1100]",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ResponseError": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/model.ResponseErrorField"
                },
                "success": {
                    "type": "boolean",
                    "default": false
                }
            }
        },
        "model.ResponseErrorField": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                }
            }
        },
        "model.ResponseSuccess": {
            "type": "object",
            "properties": {
                "data": {},
                "success": {
                    "type": "boolean",
                    "default": true
                }
            }
        },
        "requests.createCreativeSpaceRequestCoordinate": {
            "type": "object",
            "required": [
                "latitude",
                "longitude"
            ],
            "properties": {
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                }
            }
        },
        "requests.createCreativeSpaceRequestData": {
            "type": "object",
            "required": [
                "coordinate",
                "description",
                "metroStations",
                "photos",
                "pricePerHour",
                "workingHours"
            ],
            "properties": {
                "coordinate": {
                    "$ref": "#/definitions/requests.createCreativeSpaceRequestCoordinate"
                },
                "description": {
                    "type": "string"
                },
                "metroStations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/requests.createCreativeSpaceRequestMetroStation"
                    }
                },
                "photos": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "pricePerHour": {
                    "type": "integer"
                },
                "workingHours": {
                    "$ref": "#/definitions/requests.createCreativeSpaceRequestWorkingHours"
                }
            }
        },
        "requests.createCreativeSpaceRequestMetroStation": {
            "type": "object",
            "required": [
                "distanceInMinutes",
                "id"
            ],
            "properties": {
                "distanceInMinutes": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "requests.createCreativeSpaceRequestWorkingHours": {
            "type": "object",
            "required": [
                "endAt",
                "startAt"
            ],
            "properties": {
                "endAt": {
                    "type": "string"
                },
                "startAt": {
                    "type": "string"
                }
            }
        },
        "requests.createSessionRequestData": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "requests.createUserRequestData": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "requests.patchUserRequestData": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "responses.createCreativeSpaceResponseCreativeSpace": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "responses.createCreativeSpaceResponseData": {
            "type": "object",
            "properties": {
                "creativeSpace": {
                    "$ref": "#/definitions/responses.createCreativeSpaceResponseCreativeSpace"
                }
            }
        },
        "responses.createSessionResponseData": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/responses.createSessionResponseUser"
                }
            }
        },
        "responses.createSessionResponseUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "responses.createUserResponseData": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/responses.createUserResponseUser"
                }
            }
        },
        "responses.createUserResponseUser": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "responses.getMetroStationsResponseData": {
            "type": "object",
            "properties": {
                "metroStations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/responses.getMetroStationsResponseMetroStation"
                    }
                }
            }
        },
        "responses.getMetroStationsResponseMetroStation": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "responses.getUserResponseData": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/responses.getUserResponseUser"
                }
            }
        },
        "responses.getUserResponseUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "responses.getUsersResponseData": {
            "type": "object",
            "properties": {
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/responses.getUsersResponseUser"
                    }
                }
            }
        },
        "responses.getUsersResponseUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "responses.patchUserResponseData": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/responses.patchUserResponseUser"
                }
            }
        },
        "responses.patchUserResponseUser": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8080.",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Starter API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
