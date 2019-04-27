package repository

import (
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
	if err := r.db.Where(`resume_id = ?`, resumeID).Find(&experiences).Error; err != nil {
		res.Error = err
		return
	}
	res.Data = experiences
	return
}

func (r *experienceRepo) Save(data *domain.Experience) (res shared.Result) {
	var exp domain.Experience
	if err := r.db.Where(domain.Experience{Title: data.Title, Company: data.Company}).Assign(data).FirstOrCreate(&exp).Error; err != nil {
		res.Error = err
		return
	}

	res.Data = &exp
	return
}
