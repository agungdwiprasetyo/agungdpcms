package usecase

import (
	"github.com/agungdwiprasetyo/agungdpcms/src/chat"
)

// Chat abstraction
type Chat interface {
	Join(roomID string, client *chat.Client)
}
