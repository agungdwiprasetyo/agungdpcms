package delivery

import (
	"context"

	"github.com/agungdwiprasetyo/agungdpcms/src/chat/serializer"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/usecase"
)

type GetAllMessageArgs struct {
	GroupID int32
}

type GraphqlHandler struct {
	uc usecase.Chat
}

func NewGraphqlHandler(uc usecase.Chat) *GraphqlHandler {
	return &GraphqlHandler{uc}
}

func (h *GraphqlHandler) GetAllMessage(ctx context.Context, args *GetAllMessageArgs) (*serializer.MessageListSchema, error) {
	res := h.uc.FindAllMessagesByGroupID(args.GroupID)
	if res.Error != nil {
		return nil, res.Error
	}

	return res.Data.(*serializer.MessageListSchema), nil
}
