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
        "termsOfService": "http://swagger.io/terms/",
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
        "/v1/auth": {
            "get": {
                "description": "Sending login and password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Defining an authorized user",
                "responses": {
                    "200": {
                        "description": "successfully auth",
                        "schema": {
                            "$ref": "#/definitions/models.UserAuthRequest"
                        }
                    },
                    "400": {
                        "description": "return error",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.ErrResponseAuthDefault"
                        }
                    },
                    "401": {
                        "description": "no cookie",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.ErrResponseAuthNoCookie"
                        }
                    },
                    "404": {
                        "description": "such cookie not found",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.ErrResponseAuthNoSuchCookie"
                        }
                    },
                    "405": {
                        "description": "method not allowed"
                    },
                    "500": {
                        "description": "something unusual has happened"
                    }
                }
            }
        },
        "/v1/auth/login": {
            "post": {
                "description": "Sending login and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "User authentication",
                "parameters": [
                    {
                        "description": "Request body for login",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successfully login",
                        "schema": {
                            "$ref": "#/definitions/models.UserLoginResponse"
                        }
                    },
                    "400": {
                        "description": "return error",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.ErrResponseAuthDefault"
                        }
                    },
                    "404": {
                        "description": "such user not found",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.ErrResponseAuthNoSuchUser"
                        }
                    },
                    "405": {
                        "description": "method not allowed"
                    },
                    "500": {
                        "description": "something unusual has happened"
                    }
                }
            }
        },
        "/v1/auth/logout": {
            "get": {
                "description": "Session delete",
                "tags": [
                    "user"
                ],
                "summary": "User logout",
                "responses": {
                    "204": {
                        "description": "successfully logout"
                    },
                    "400": {
                        "description": "return error",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.ErrResponseAuthDefault"
                        }
                    },
                    "401": {
                        "description": "no cookie",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.ErrResponseAuthNoCookie"
                        }
                    },
                    "404": {
                        "description": "such cookie not found",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.ErrResponseAuthNoSuchCookie"
                        }
                    },
                    "405": {
                        "description": "method not allowed"
                    },
                    "500": {
                        "description": "something unusual has happened"
                    }
                }
            }
        },
        "/v1/auth/signup": {
            "post": {
                "description": "Sending login and password for registration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "New user registration",
                "parameters": [
                    {
                        "description": "Request body for login",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserSignupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successfully login",
                        "schema": {
                            "$ref": "#/definitions/models.UserSignupResponse"
                        }
                    },
                    "400": {
                        "description": "return error",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.ErrResponseAuthDefault"
                        }
                    },
                    "404": {
                        "description": "such user not found",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.ErrResponseAuthNoSuchUser"
                        }
                    },
                    "405": {
                        "description": "method not allowed"
                    },
                    "500": {
                        "description": "something unusual has happened"
                    }
                }
            }
        },
        "/v1/popular_films": {
            "get": {
                "description": "Films from the \"popular\" category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "collections"
                ],
                "summary": "Popular movies",
                "responses": {
                    "200": {
                        "description": "returns an array of movies",
                        "schema": {
                            "$ref": "#/definitions/models.FilmCollectionRequest"
                        }
                    },
                    "400": {
                        "description": "return error"
                    },
                    "405": {
                        "description": "method not allowed"
                    },
                    "500": {
                        "description": "something unusual has happened"
                    }
                }
            }
        }
    },
    "definitions": {
        "httpmodels.ErrResponseAuthDefault": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Auth: [{{Reason}}]"
                }
            }
        },
        "httpmodels.ErrResponseAuthNoCookie": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Auth: [request has no cookies]"
                }
            }
        },
        "httpmodels.ErrResponseAuthNoSuchCookie": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Auth: [no such cookie]"
                }
            }
        },
        "httpmodels.ErrResponseAuthNoSuchUser": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Auth: [such user doesn't exist]"
                }
            }
        },
        "models.FilmCollectionRequest": {
            "type": "object",
            "properties": {
                "films": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.filmInCollectionRequest"
                    }
                },
                "title": {
                    "type": "string",
                    "example": "Популярное"
                }
            }
        },
        "models.UserAuthRequest": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string",
                    "example": "{{ссылка}}"
                },
                "email": {
                    "type": "string",
                    "example": "dop123@mail.ru"
                },
                "nickname": {
                    "type": "string",
                    "example": "Bot373"
                }
            }
        },
        "models.UserLoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "YasaPupkinEzji@top.world"
                },
                "nickname": {
                    "type": "string",
                    "example": "StepByyyy"
                },
                "password": {
                    "type": "string",
                    "example": "Widget Adapter"
                }
            }
        },
        "models.UserLoginResponse": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string",
                    "example": "{{ссылка}}"
                },
                "email": {
                    "type": "string",
                    "example": "dop123@mail.ru"
                },
                "nickname": {
                    "type": "string",
                    "example": "StepByyyy"
                }
            }
        },
        "models.UserSignupRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "YasaPupkinEzji@top.world"
                },
                "nickname": {
                    "type": "string",
                    "example": "StepByyyy"
                },
                "password": {
                    "type": "string",
                    "example": "Widget Adapter"
                }
            }
        },
        "models.UserSignupResponse": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string",
                    "example": "{{ссылка}}"
                },
                "email": {
                    "type": "string",
                    "example": "dop123@mail.ru"
                },
                "nickname": {
                    "type": "string",
                    "example": "StepByyyy"
                }
            }
        },
        "models.filmInCollectionRequest": {
            "type": "object",
            "properties": {
                "film_id": {
                    "type": "integer",
                    "example": 23
                },
                "film_name": {
                    "type": "string",
                    "example": "Game of Thrones"
                },
                "genres": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "фэнтези",
                        " приключения"
                    ]
                },
                "poster_ver": {
                    "type": "string",
                    "example": "{{ссылка}}"
                },
                "ratio": {
                    "type": "string",
                    "example": "7.9"
                },
                "year_prod": {
                    "type": "string",
                    "example": "2014"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "movie-gate.online",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MovieGate",
	Description:      "Server for MovieGate application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
