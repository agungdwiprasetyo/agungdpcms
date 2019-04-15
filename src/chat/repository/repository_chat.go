package repository

import (
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/domain"
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/jinzhu/gorm"
)

type chatGorm struct {
	db *gorm.DB
}

// NewChatRepo gorm
func NewChatRepo(db *gorm.DB) Chat {
	return &chatGorm{db}
}

func (r *chatGorm) FindGroupByID(id int) (res shared.Result) {
	var group domain.Group
	if err := r.db.First(&group, id).Error; err != nil {
		res.Error = err
	}
	return
}

func (r *chatGorm) FindAllMessageByGroupID(groupID int) (res shared.Result) {
	var messages []*domain.Message
	if err := r.db.Model(&domain.Group{ID: groupID}).Related(&messages).Error; err != nil {
		res.Error = err
	}
	res.Data = messages
	return res
}

func (r *chatGorm) SaveMessage(data *domain.Message) (res shared.Result) {
	res.Data = data
	if err := r.db.Save(data).Error; err != nil {
		res.Error = err
	}
	return
}
