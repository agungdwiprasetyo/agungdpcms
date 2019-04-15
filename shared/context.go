package shared

import (
	"context"
	"net/http"
)

// ContextKey type
type ContextKey string

// ParseHeaderFromContext extract http.Header from http context
func ParseHeaderFromContext(ctx context.Context) http.Header {
	ctxVal := ctx.Value(ContextKey("headers"))
	headers, ok := ctxVal.(http.Header)
	if !ok {
		return http.Header{}
	}

	return headers
}

// ParseUserData extract userdata from given context
func ParseUserData(ctx context.Context) interface{} {
	return ctx.Value(ContextKey("userData"))
}
