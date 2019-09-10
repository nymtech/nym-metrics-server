// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-09-10 17:53:21.332726396 +0100 BST m=+0.080374035

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a temporarily centralized directory/PKI/metrics API to allow us to get the other Nym node types running. Its functionality will eventually be folder into other parts of Nym.",
        "title": "Nym Directory API",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "https://github.com/nymtech/directory-server/license"
        },
        "version": "1.0"
    },
    "host": "directory.nymtech.net",
    "basePath": "/v1",
    "paths": {
        "/api/metrics": {
            "post": {
                "description": "You'd never want to run this in production, but for demo and debug purposes it gives us the ability to generate useful visualisations of network traffic.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "metrics"
                ],
                "summary": "Create a metric detailing how many messages a given mixnode sent",
                "operationId": "createMetric",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ObjectRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ObjectIDResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.Error"
                        }
                    }
                }
            }
        },
        "/api/nodes": {
            "post": {
                "description": "TODO",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pki"
                ],
                "summary": "Create a node in the PKI",
                "operationId": "createObject",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ObjectRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.ObjectIDResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/rest.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "rest.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "rest.ObjectIDResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "rest.ObjectRequest": {
            "type": "object",
            "properties": {
                "object": {
                    "type": "object"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
