package repository

import (
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/jinzhu/gorm"
)

type achievementRepo struct {
	db *gorm.DB
}

// NewAchievementRepository construct new resume repo
func NewAchievementRepository(db *gorm.DB) Achievement {
	return &achievementRepo{db}
}

func (r *achievementRepo) FindByResumeID(resumeID int) *shared.Result {
	var achievements []*domain.Achievement

	if err := r.db.Where(`resume_id = ?`, resumeID).Find(&achievements).Error; err != nil {
		return &shared.Result{Error: err}
	}

	return &shared.Result{Data: achievements}
}
