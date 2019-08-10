package repository

import (
	"fmt"

	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/jinzhu/gorm"
)

type achievementRepo struct {
	db *gorm.DB
}

// NewAchievementRepository construct new resume repo
func NewAchievementRepository(db *gorm.DB) Achievement {
	db.AutoMigrate(&domain.Achievement{})
	return &achievementRepo{db}
}

func (r *achievementRepo) FindByResumeID(resumeID int) <-chan []*domain.Achievement {
	output := make(chan []*domain.Achievement)

	go func() {
		defer close(output)

		var achievements []*domain.Achievement
		if err := r.db.Where(domain.Achievement{ResumeID: resumeID}).Find(&achievements).Error; err != nil {
			panic(err)
		}

		output <- achievements
	}()

	return output
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
	db := r.db.Delete(data)
	if err := db.Error; err != nil {
		res.Error = err
	}
	if affected := db.RowsAffected; affected == 0 {
		res.Error = fmt.Errorf("data with id=%d not found", data.ID)
	}
	return
}
