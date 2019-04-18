package repository

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
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

func (r *achievementRepo) Save(data *domain.Achievement) shared.Result {
	var ach domain.Achievement
	if err := r.db.Where(domain.Achievement{Name: data.Name}).Assign(data).FirstOrCreate(&ach).Error; err != nil {
		return shared.Result{Error: err}
	}

	return shared.Result{Data: &ach}
}
