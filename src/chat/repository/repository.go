package repository

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/domain"
	"github.com/jinzhu/gorm"
)

type (
	// Chat abstraction
	Chat interface {
		FindGroupByID(id int) shared.Result
		FindAllMessage(groupID, offset, limit int) shared.Result
		CountByGroupID(groupID int) shared.Result
		SaveMessage(data *domain.Message) shared.Result
	}
)

// Repository model
type Repository struct {
	Chat Chat
}

// NewRepository constructor
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Chat: NewChatRepo(db),
	}
}
