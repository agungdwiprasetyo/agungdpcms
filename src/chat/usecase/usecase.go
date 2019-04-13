package usecase

import (
	"github.com/agungdwiprasetyo/agungdpcms/src/chat"
	"github.com/agungdwiprasetyo/agungdpcms/src/shared"
)

// Chat abstraction
type Chat interface {
	Join(roomID string, client *chat.Client) error
	FindAllMessagesByGroupID(groupID int32) shared.Result
}
