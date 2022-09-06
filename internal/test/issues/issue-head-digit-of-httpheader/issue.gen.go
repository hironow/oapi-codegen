// Package head_digit_of_httpheader provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package head_digit_of_httpheader

import (
	"context"
)

type N200ResponseHeaders struct {
	N000Foo string
}
type N200Response struct {
	Headers N200ResponseHeaders
}

type GetFooRequestObject struct {
}

type GetFoo200Response = N200Response

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /foo)
	GetFoo(ctx context.Context, request GetFooRequestObject) interface{}
}