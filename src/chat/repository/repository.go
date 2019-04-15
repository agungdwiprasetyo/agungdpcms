package repository

import (
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/domain"
	"github.com/agungdwiprasetyo/agungdpcms/shared"
)

// Chat abstraction
type Chat interface {
	FindGroupByID(id int) shared.Result
	FindAllMessageByGroupID(groupID int) shared.Result
	SaveMessage(data *domain.Message) shared.Result
}
