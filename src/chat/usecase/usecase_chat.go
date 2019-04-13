package usecase

import (
	"github.com/agungdwiprasetyo/agungdpcms/src/chat"
)

type chatImpl struct{}

// New chat usecase
func New() Chat {
	return &chatImpl{}
}

func (uc *chatImpl) Join(roomID string, client *chat.Client) {
	// TODO: add client join to roomId in database (repository)

	// Read client message
	client.Read()
	// Write client message
	client.Write()
}
