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

func (r *achievementRepo) FindByResumeID(resumeID int) (res shared.Result) {
	var achievements []*domain.Achievement

	if err := r.db.Where(domain.Achievement{ResumeID: resumeID}).Find(&achievements).Error; err != nil {
		res.Error = err
		return
	}

	res.Data = achievements
	return
}

func (r *achievementRepo) Save(data *domain.Achievement) shared.Result {
	var ach domain.Achievement
	where := domain.Achievement{
		Name: data.Name, Appreciator: data.Appreciator, Year: data.Year,
	}
	if err := r.db.Where(where).Assign(data).FirstOrCreate(&ach).Error; err != nil {
		return shared.Result{Error: err}
	}

	return shared.Result{Data: &ach}
}

func (r *achievementRepo) Remove(data *domain.Achievement) (res shared.Result) {
	if err := r.db.Delete(data).Error; err != nil {
		res.Error = err
	}
	return
}
