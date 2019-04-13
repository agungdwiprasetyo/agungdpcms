package middleware

import (
	"context"

	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/src/shared"
	"github.com/agungdwiprasetyo/go-utils/debug"
)

type bearer struct {
	conf *config.Config
}

// NewBearer constructor
func NewBearer(conf *config.Config) Middleware {
	return &bearer{conf}
}

func (b *bearer) WithAuth(ctx context.Context) {
	headers := shared.ParseHeaderFromContext(ctx)
	debug.Println(headers.Get("Authorization"))
	panic(headers.Get("Authorization"))
}
