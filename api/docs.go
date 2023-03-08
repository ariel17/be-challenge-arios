// Package api GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Ariel Gerardo Ríos",
            "url": "http://ariel17.com.ar",
            "email": "arielgerardorios@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/competitions/:code/players": {
            "get": {
                "description": "Given the competition code, if it exists on database, returns all players from all participating teams.",
                "produces": [
                    "application/json"
                ],
                "summary": "Shows all players from a given competition.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Competition code to filter players.",
                        "name": "code",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Team name to filter players by",
                        "name": "teamName",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.PlayersResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/server.PlayersResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.PlayersResult"
                        }
                    }
                }
            }
        },
        "/importer": {
            "post": {
                "description": "Enqueues data scrapping from football-data.org API based on competition code. It is a background process so this endpoint only reflects the state of the petition.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Imports football data by competition code.",
                "parameters": [
                    {
                        "description": "Competition code to import.",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.ImporterCommand"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/server.Status"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/server.Status"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.Status"
                        }
                    }
                }
            }
        },
        "/status": {
            "get": {
                "description": "Returns a JSON reflecting the application's health.",
                "produces": [
                    "application/json"
                ],
                "summary": "Shows the status of the application.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.Status"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.Status"
                        }
                    }
                }
            }
        },
        "/teams/:tla": {
            "get": {
                "description": "If indicated players/coach also can be resolved if they exist.",
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves indicated team details with players/coach.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Team TLA value to fetch.",
                        "name": "tla",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Resolve team players/coach if present.",
                        "name": "showPlayers",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.TeamResult"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.TeamResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.TeamResult"
                        }
                    }
                }
            }
        },
        "/teams/:tla/persons": {
            "get": {
                "description": "Retrieves all persons on a team (players/coach).",
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves all persons on a team (players/coach).",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Team TLA value to fetch.",
                        "name": "tla",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.PersonsResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.PersonsResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Person": {
            "type": "object",
            "properties": {
                "date_of_birth": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nationality": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                }
            }
        },
        "models.Team": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "area_name": {
                    "type": "string"
                },
                "coach": {
                    "$ref": "#/definitions/models.Person"
                },
                "name": {
                    "type": "string"
                },
                "players": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Person"
                    }
                },
                "short_name": {
                    "type": "string"
                },
                "tla": {
                    "type": "string"
                }
            }
        },
        "server.ImporterCommand": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        },
        "server.PersonsResult": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string"
                },
                "ok": {
                    "type": "boolean"
                },
                "persons": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Person"
                    }
                }
            }
        },
        "server.PlayersResult": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string"
                },
                "ok": {
                    "type": "boolean"
                },
                "players": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Person"
                    }
                }
            }
        },
        "server.Status": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string"
                },
                "ok": {
                    "type": "boolean"
                }
            }
        },
        "server.TeamResult": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string"
                },
                "ok": {
                    "type": "boolean"
                },
                "team": {
                    "$ref": "#/definitions/models.Team"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "BE Challenge by Ariel Gerardo Ríos",
	Description:      "A challenge that uses football-data.org data on its own models.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
