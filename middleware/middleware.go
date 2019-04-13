package middleware

import (
	"context"
)

// Middleware abstraction
type Middleware interface {
	WithAuth(ctx context.Context)
}
