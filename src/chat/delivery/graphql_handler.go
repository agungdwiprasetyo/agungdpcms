package delivery

import (
	"context"

	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/serializer"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/usecase"
	"github.com/agungdwiprasetyo/go-utils/debug"
)

// GraphqlHandler model
type GraphqlHandler struct {
	uc   usecase.Chat
	midd middleware.Middleware
}

// NewGraphqlHandler constructor
func NewGraphqlHandler(uc usecase.Chat, midd middleware.Middleware) *GraphqlHandler {
	return &GraphqlHandler{
		uc:   uc,
		midd: midd,
	}
}

// GetAllMessage graphql handler
func (h *GraphqlHandler) GetAllMessage(ctx context.Context, args *GetAllMessageArgs) (*serializer.MessageListSchema, error) {
	ctx = h.midd.WithAuth(ctx)
	userData := shared.ParseUserData(ctx)
	debug.PrintJSON(userData)

	res := h.uc.FindAllMessagesByGroupID(args.GroupID)
	if res.Error != nil {
		return nil, res.Error
	}

	return res.Data.(*serializer.MessageListSchema), nil
}
