package delivery

import (
	"context"

	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	"github.com/agungdwiprasetyo/agungdpcms/src/user/serializer"
	"github.com/agungdwiprasetyo/agungdpcms/src/user/usecase"
)

// GraphQLHandler model
type GraphQLHandler struct {
	uc   usecase.User
	midd middleware.Middleware
}

// NewGraphQLHandler delivery
func NewGraphQLHandler(uc usecase.User, midd middleware.Middleware) *GraphQLHandler {
	return &GraphQLHandler{
		uc:   uc,
		midd: midd,
	}
}

// Login handler
func (h *GraphQLHandler) Login(ctx context.Context, args *LoginArgs) (*serializer.UserSchema, error) {
	res := h.uc.Login(args.Username, args.Password)
	if res.Error != nil {
		return nil, res.Error
	}

	data := res.Data.(*serializer.UserSchema)
	return data, nil
}
