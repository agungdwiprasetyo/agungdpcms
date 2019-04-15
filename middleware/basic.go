package middleware

import (
	"context"
	"encoding/base64"
	"strings"

	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/helper"
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/go-utils/debug"
)

type basic struct {
	username, password string
}

// NewBasicAuth construct new basic auth middleware
func NewBasicAuth(conf *config.Config) Middleware {
	return &basic{username: conf.Env.Username, password: conf.Env.Password}
}

func (b *basic) WithAuth(ctx context.Context) context.Context {
	headers := shared.ParseHeaderFromContext(ctx)
	authorizations := strings.Split(headers.Get("Authorization"), " ")
	if len(authorizations) != 2 {
		panic("Invalid authorization")
	}

	tp, val := authorizations[0], authorizations[1]
	if tp != helper.Basic {
		panic("Invalid authorization")
	}

	data, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		panic(err)
	}

	decoded := strings.Split(string(data), ":")
	username, password := decoded[0], decoded[1]

	if username != b.username || password != b.password {
		panic("Forbidden")
	}

	debug.Println(username, password)

	return ctx
}
