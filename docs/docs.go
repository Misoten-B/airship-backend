// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ar_assets/{ar_assets_id}": {
            "get": {
                "tags": [
                    "ArAssets"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ArAssets ID",
                        "name": "ar_assets_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.ArAssetsResponse"
                        }
                    }
                }
            }
        },
        "/business_card_parts_coordinate": {
            "get": {
                "tags": [
                    "BusinessCardPartsCoordinate"
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.BusinessCardPartsCoordinateResponse"
                            }
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "User ID",
                        "name": "CreateUserRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResponse"
                        }
                    }
                }
            }
        },
        "/user/ar_assets": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "ArAssets"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.ArAssetsResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "ArAssets"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Image file to be uploaded",
                        "name": "qrcodeImage",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "description": "ArAssets",
                        "name": "dto.CreateArAssetsRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateArAssetsRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/user/ar_assets/{ar_assets_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "ArAssets"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ArAssets ID",
                        "name": "ar_assets_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.ArAssetsResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "ArAssets"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "ArAssets",
                        "name": "dto.CreateArAssetsRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateArAssetsRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "ArAssets ID",
                        "name": "ar_assets_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Image file to be uploaded",
                        "name": "qrcodeIcon",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "ArAssets"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ArAssets ID",
                        "name": "ar_assets_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/user/business_card": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "BusinessCard"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.BusinessCardResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "BusinessCard"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Image file to be uploaded",
                        "name": "BusinessCardBackgroundImage",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "description": "BusinessCard",
                        "name": "CreateBusinessCardRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateBusinessCardRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/user/business_card/{business_card_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "BusinessCard"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "BusinessCard ID",
                        "name": "business_card_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.BusinessCardResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "BusinessCard"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "BusinessCard ID",
                        "name": "business_card_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Image file to be uploaded",
                        "name": "BusinessCardBackgroundImage",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "description": "BusinessCard",
                        "name": "CreateBusinessCardRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateBusinessCardRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "BusinessCard"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "BusinessCard ID",
                        "name": "business_card_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/user/business_card_background": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "BusinessCardBackground"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.BusinessCardBackgroundResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "BusinessCardBackground"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "BusinessCardBackground",
                        "name": "dto.CreateBackgroundRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateBackgroundRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/user/three_dimentional": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "ThreeDimentionalModel"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.ThreeDimentionalResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "ThreeDimentionalModel"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "3dmodel file to be uploaded",
                        "name": "ThreeDimentionalModel",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.ThreeDimentionalResponse"
                        }
                    }
                }
            }
        },
        "/user/three_dimentional/{three_dimentional_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "ThreeDimentionalModel"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ThreeDimentional ID",
                        "name": "three_dimentional_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ThreeDimentionalResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "ThreeDimentionalModel"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ThreeDimentional ID",
                        "name": "three_dimentional_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "3dmodel file to be uploaded",
                        "name": "ThreeDimentionalModel",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.ThreeDimentionalResponse"
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
                "tags": [
                    "ThreeDimentionalModel"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ThreeDimentional ID",
                        "name": "three_dimentional_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/user/{user_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "update user",
                        "name": "CreateUserRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer [Firebase JWT Token]",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ArAssetsResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "id"
                },
                "qrcode_image_path": {
                    "type": "string",
                    "example": "url"
                },
                "speaking_audio_path": {
                    "type": "string",
                    "example": "url"
                },
                "speaking_description": {
                    "type": "string",
                    "example": "description"
                },
                "three_dimentional_path": {
                    "type": "string",
                    "example": "url"
                }
            }
        },
        "dto.BusinessCardBackgroundResponse": {
            "type": "object",
            "properties": {
                "business_card_background_color": {
                    "type": "string",
                    "example": "#ffffff"
                },
                "business_card_background_image": {
                    "type": "string",
                    "example": "url"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "dto.BusinessCardPartsCoordinateResponse": {
            "type": "object",
            "properties": {
                "address_x": {
                    "type": "integer"
                },
                "address_y": {
                    "type": "integer"
                },
                "company_name_x": {
                    "type": "integer"
                },
                "company_name_y": {
                    "type": "integer"
                },
                "department_x": {
                    "type": "integer"
                },
                "department_y": {
                    "type": "integer"
                },
                "display_name_x": {
                    "type": "integer"
                },
                "display_name_y": {
                    "type": "integer"
                },
                "email_x": {
                    "type": "integer"
                },
                "email_y": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "official_position_x": {
                    "type": "integer"
                },
                "official_position_y": {
                    "type": "integer"
                },
                "phone_number_x": {
                    "type": "integer"
                },
                "phone_number_y": {
                    "type": "integer"
                },
                "postal_code_x": {
                    "type": "integer"
                },
                "postal_code_y": {
                    "type": "integer"
                },
                "qrcode_x": {
                    "type": "integer"
                },
                "qrcode_y": {
                    "type": "integer"
                }
            }
        },
        "dto.BusinessCardResponse": {
            "type": "object",
            "properties": {
                "accessCount": {
                    "type": "integer"
                },
                "address": {
                    "type": "string"
                },
                "businessCardBackgroundColor": {
                    "description": "background",
                    "type": "string",
                    "example": "#ffffff"
                },
                "businessCardBackgroundImage": {
                    "type": "string",
                    "example": "url"
                },
                "businessCardName": {
                    "type": "string"
                },
                "businessCardPartsCoordinate": {
                    "description": "business card",
                    "type": "string"
                },
                "companyName": {
                    "type": "string"
                },
                "department": {
                    "type": "string"
                },
                "displayName": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "officialPosition": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "postalCode": {
                    "type": "string"
                },
                "speakingAudioPath": {
                    "type": "string"
                },
                "speakingDescription": {
                    "type": "string"
                },
                "threeDimentionalModel": {
                    "description": "ar assets",
                    "type": "string"
                }
            }
        },
        "dto.CreateArAssetsRequest": {
            "type": "object",
            "properties": {
                "speaking_description": {
                    "type": "string",
                    "example": "description"
                },
                "three_dimentional_ID": {
                    "type": "string",
                    "example": "url"
                }
            }
        },
        "dto.CreateBackgroundRequest": {
            "type": "object",
            "properties": {
                "business_card_background": {
                    "type": "string",
                    "example": "#ffffff"
                },
                "business_card_background_image": {
                    "type": "string",
                    "x-nullable": true,
                    "example": "url"
                }
            }
        },
        "dto.CreateBusinessCardRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "x-nullable": true,
                    "example": "東京都渋谷区神南1-1-1"
                },
                "ar_assets_id": {
                    "description": "ar assets",
                    "type": "string",
                    "x-nullable": true,
                    "example": "ar_assets_id"
                },
                "business_card_background_id": {
                    "description": "background",
                    "type": "string",
                    "x-nullable": true,
                    "example": "id"
                },
                "business_card_name": {
                    "type": "string",
                    "x-nullable": true,
                    "example": "会社"
                },
                "business_card_parts_coordinate": {
                    "description": "business card",
                    "type": "string",
                    "x-nullable": true,
                    "example": "id"
                },
                "company_name": {
                    "type": "string",
                    "x-nullable": true,
                    "example": "会社名"
                },
                "department": {
                    "type": "string",
                    "x-nullable": true,
                    "example": "部署"
                },
                "display_name": {
                    "type": "string",
                    "x-nullable": true,
                    "example": "名前"
                },
                "email": {
                    "type": "string",
                    "x-nullable": true,
                    "example": "sample@example.com"
                },
                "official_position": {
                    "type": "string",
                    "x-nullable": true,
                    "example": "役職"
                },
                "phone_number": {
                    "type": "string",
                    "x-nullable": true,
                    "example": "090-1234-5678"
                },
                "postal_code": {
                    "type": "string",
                    "x-nullable": true,
                    "example": "123-4567"
                },
                "qrcode_image_path": {
                    "type": "string",
                    "x-nullable": true,
                    "example": "url"
                }
            }
        },
        "dto.CreateUserRequest": {
            "type": "object",
            "properties": {
                "is_toured": {
                    "type": "boolean",
                    "x-nullable": true,
                    "example": false
                }
            }
        },
        "dto.ThreeDimentionalResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "dto.UserResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "is_toured": {
                    "type": "boolean"
                },
                "recorded_model_path": {
                    "type": "string"
                },
                "recorded_voice_path": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "AIRship API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
