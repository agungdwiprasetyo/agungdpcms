package repository

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/domain"
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

func (r *chatGorm) FindAllMessage(groupID, offset, limit int) (res shared.Result) {
	var messages []*domain.Message
	if err := r.db.Model(&domain.Group{ID: groupID}).Limit(limit).Offset(offset).Related(&messages).Error; err != nil {
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

func (r *chatGorm) CountByGroupID(groupID int) (res shared.Result) {
	var count int
	if err := r.db.Model(&domain.Message{GroupID: groupID}).Count(&count).Error; err != nil {
		res.Error = err
	}
	res.Data = count
	return res
}
