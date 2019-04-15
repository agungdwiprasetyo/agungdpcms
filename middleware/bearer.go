package middleware

import (
	"context"
	"crypto/rsa"
	"strings"

	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/shared/token"
)

type bearer struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	token      token.Token
}

// NewBearer constructor
func NewBearer(conf *config.Config, token token.Token) Middleware {
	return &bearer{
		privateKey: conf.PrivateKey,
		publicKey:  conf.PublicKey,
		token:      token,
	}
}

func (b *bearer) WithAuth(ctx context.Context) context.Context {
	headers := shared.ParseHeaderFromContext(ctx)

	authorization := headers.Get("Authorization")
	if authorization == "" {
		panic("Invalid authorization")
	}

	authValues := strings.Split(authorization, " ")
	if strings.ToLower(authValues[0]) != "bearer" || len(authValues) != 2 {
		panic("Invalid authorization")
	}

	tokenString := authValues[1]
	claims, ok := b.token.Extract(tokenString)
	if !ok {
		panic("Invalid authorization")
	}

	return context.WithValue(ctx, shared.ContextKey("userData"), claims)
}
