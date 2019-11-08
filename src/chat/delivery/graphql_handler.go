package delivery

import (
	"context"

	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/domain"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/serializer"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/usecase"
	"github.com/agungdwiprasetyo/go-utils/debug"
)

// GraphQLHandler model
type GraphQLHandler struct {
	uc   usecase.Chat
	midd middleware.Middleware
}

// NewGraphQLHandler constructor
func NewGraphQLHandler(uc usecase.Chat, midd middleware.Middleware) *GraphQLHandler {
	return &GraphQLHandler{
		uc:   uc,
		midd: midd,
	}
}

// GetAllMessage graphql handler
func (h *GraphQLHandler) GetAllMessage(ctx context.Context, args *domain.Param) (*serializer.MessageListSchema, error) {
	ctx = h.midd.WithAuth(ctx)
	userData := shared.ParseUserData(ctx)
	debug.PrintJSON(userData)

	res := h.uc.FindAllMessages(args)
	if res.Error != nil {
		return nil, res.Error
	}

	return res.Data.(*serializer.MessageListSchema), nil
}
