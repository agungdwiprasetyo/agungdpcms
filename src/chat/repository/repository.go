package repository

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/domain"
)

// Chat abstraction
type Chat interface {
	FindGroupByID(id int) shared.Result
	FindAllMessage(groupID, offset, limit int) shared.Result
	CountByGroupID(groupID int) shared.Result
	SaveMessage(data *domain.Message) shared.Result
}
