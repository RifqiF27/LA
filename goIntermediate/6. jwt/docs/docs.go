// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://example.com/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://example.com/support",
            "email": "support@example.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "Authenticate user with email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Login request payload",
                        "name": "loginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful login",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.ResponseOK"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/utils.LoginResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid email or password",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to save token",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register a new user with email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration request payload",
                        "name": "registerRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User registered successfully",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseOK"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to register user",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/vouchers": {
            "get": {
                "security": [
                    {
                        "Authentication": []
                    },
                    {
                        "UserID": []
                    }
                ],
                "description": "Retrieve vouchers based on status, area, and voucher type",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vouchers"
                ],
                "summary": "Get vouchers by query parameters",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Voucher status",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Voucher area",
                        "name": "area",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Voucher type",
                        "name": "voucher_type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Vouchers retrieved successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.ResponseOK"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Voucher"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Vouchers not found",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/vouchers/create": {
            "post": {
                "security": [
                    {
                        "Authentication": []
                    },
                    {
                        "UserID": []
                    }
                ],
                "description": "Create a new voucher with provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vouchers"
                ],
                "summary": "Create a new voucher",
                "parameters": [
                    {
                        "description": "Voucher details",
                        "name": "voucher",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Voucher"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Created successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.ResponseOK"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Voucher"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to create voucher",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/vouchers/redeem": {
            "post": {
                "security": [
                    {
                        "Authentication": []
                    },
                    {
                        "UserID": []
                    }
                ],
                "description": "Redeem a voucher using points",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vouchers"
                ],
                "summary": "Create a redeem voucher",
                "parameters": [
                    {
                        "description": "Redeem request payload",
                        "name": "redeemRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/managementvoucherhandler.RedeemRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Redeem created successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.ResponseOK"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Redeem"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to create redeem voucher",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/vouchers/redeem-points": {
            "get": {
                "security": [
                    {
                        "Authentication": []
                    },
                    {
                        "UserID": []
                    }
                ],
                "description": "Retrieve the list of redeem points",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vouchers"
                ],
                "summary": "Show redeem points",
                "responses": {
                    "200": {
                        "description": "Redeem points retrieved successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.ResponseOK"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Redeem"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Redeem points not found",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/vouchers/{id}": {
            "put": {
                "security": [
                    {
                        "Authentication": []
                    },
                    {
                        "UserID": []
                    }
                ],
                "description": "Update a voucher by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vouchers"
                ],
                "summary": "Update a voucher",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Voucher ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated voucher details",
                        "name": "voucher",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Voucher"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.ResponseOK"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Voucher"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to update voucher",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Authentication": []
                    },
                    {
                        "UserID": []
                    }
                ],
                "description": "Soft delete a voucher by ID",
                "tags": [
                    "Vouchers"
                ],
                "summary": "Soft delete a voucher",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Voucher ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseOK"
                        }
                    },
                    "500": {
                        "description": "Failed to delete voucher",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "john.doe@example.com"
                },
                "password": {
                    "type": "string",
                    "example": "password1234"
                }
            }
        },
        "managementvoucherhandler.RedeemRequest": {
            "type": "object",
            "required": [
                "points",
                "user_id",
                "voucher_id"
            ],
            "properties": {
                "points": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                },
                "voucher_id": {
                    "type": "integer"
                }
            }
        },
        "models.Redeem": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "redeem_date": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "voucher_id": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "models.Voucher": {
            "type": "object",
            "required": [
                "applicable_areas",
                "discount_value",
                "end_date",
                "minimum_purchase",
                "payment_methods",
                "quota",
                "start_date",
                "voucher_category",
                "voucher_code",
                "voucher_name",
                "voucher_type"
            ],
            "properties": {
                "applicable_areas": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Jawa"
                    ]
                },
                "description": {
                    "type": "string",
                    "example": "10% off for purchases above 200.000"
                },
                "discount_value": {
                    "type": "number",
                    "example": 10
                },
                "end_date": {
                    "type": "string",
                    "example": "2024-12-07T00:00:00Z"
                },
                "minimum_purchase": {
                    "type": "number",
                    "example": 200000
                },
                "payment_methods": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Credit Card"
                    ]
                },
                "points_required": {
                    "type": "integer",
                    "example": 220
                },
                "quota": {
                    "type": "integer",
                    "example": 50
                },
                "start_date": {
                    "type": "string",
                    "example": "2024-12-01T00:00:00Z"
                },
                "status": {
                    "type": "boolean",
                    "example": true
                },
                "voucher_category": {
                    "type": "string",
                    "example": "Free Shipping"
                },
                "voucher_code": {
                    "type": "string",
                    "example": "DESCERIA100"
                },
                "voucher_name": {
                    "type": "string",
                    "example": "PROMO GAJIAN"
                },
                "voucher_type": {
                    "type": "string",
                    "example": "redeem points"
                }
            }
        },
        "utils.ErrorResponse": {
            "type": "object",
            "properties": {
                "error_msg": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "utils.LoginResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "utils.ResponseOK": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Authentication": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "UserID": {
            "type": "apiKey",
            "name": "User-ID",
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
	Title:            "Voucher System API",
	Description:      "API for managing vouchers",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
