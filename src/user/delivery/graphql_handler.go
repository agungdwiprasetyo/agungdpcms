package delivery

import (
	"context"

	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	"github.com/agungdwiprasetyo/agungdpcms/src/user/serializer"
	"github.com/agungdwiprasetyo/agungdpcms/src/user/usecase"
)

// GraphqlHandler model
type GraphqlHandler struct {
	uc   usecase.User
	midd middleware.Middleware
}

// NewGraphqlHandler delivery
func NewGraphqlHandler(uc usecase.User, midd middleware.Middleware) *GraphqlHandler {
	return &GraphqlHandler{
		uc:   uc,
		midd: midd,
	}
}

// Login handler
func (h *GraphqlHandler) Login(ctx context.Context, args *LoginArgs) (*serializer.UserSchema, error) {
	res := h.uc.Login(args.Username, args.Password)
	if res.Error != nil {
		return nil, res.Error
	}

	data := res.Data.(*serializer.UserSchema)
	return data, nil
}
