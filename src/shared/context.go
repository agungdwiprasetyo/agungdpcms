package shared

import (
	"context"
	"net/http"
)

type ContextKey string

func ParseHeaderFromContext(ctx context.Context) http.Header {
	ctxVal := ctx.Value(ContextKey("headers"))
	headers, ok := ctxVal.(http.Header)
	if !ok {
		return http.Header{}
	}

	return headers
}
