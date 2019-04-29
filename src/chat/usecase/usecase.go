package usecase

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/domain"
	"github.com/agungdwiprasetyo/agungdpcms/websocket"
)

// Chat abstraction
type Chat interface {
	Join(roomID string, client *websocket.Client) error
	FindAllMessagesByGroupID(args *domain.GetAllMessageArgs) shared.Result
}
