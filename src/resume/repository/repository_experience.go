package repository

import (
	"fmt"

	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/jinzhu/gorm"
)

type experienceRepo struct {
	db *gorm.DB
}

// NewExperienceRepository construct new resume repo
func NewExperienceRepository(db *gorm.DB) Experience {
	return &experienceRepo{db}
}

func (r *experienceRepo) FindByResumeID(resumeID int) (res shared.Result) {
	var experiences []*domain.Experience
	if err := r.db.Where(domain.Experience{ResumeID: resumeID}).Find(&experiences).Error; err != nil {
		res.Error = err
		return
	}
	res.Data = experiences
	return
}

func (r *experienceRepo) Save(data *domain.Experience) (res shared.Result) {
	var exp domain.Experience
	where := domain.Experience{Title: data.Title, Company: data.Company}
	if err := r.db.Where(where).Assign(data).FirstOrCreate(&exp).Error; err != nil {
		res.Error = err
		return
	}

	res.Data = &exp
	return
}

func (r *experienceRepo) Remove(data *domain.Experience) (res shared.Result) {
	db := r.db.Delete(data)
	if err := db.Error; err != nil {
		res.Error = err
	}
	if affected := db.RowsAffected; affected == 0 {
		res.Error = fmt.Errorf("data with id=%d not found", data.ID)
	}
	return
}
