package usecase

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/domain"
)

// Chat abstraction
type Chat interface {
	Join(roomID string, client *chat.Client) error
	FindAllMessagesByGroupID(args *domain.GetAllMessageArgs) shared.Result
}
