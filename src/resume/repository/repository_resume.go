package repository

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/jinzhu/gorm"
)

type resumeRepo struct {
	db *gorm.DB
}

// NewResumeRepository construct new resume repo
func NewResumeRepository(db *gorm.DB) Resume {
	return &resumeRepo{db}
}

func (r *resumeRepo) FindAll() *shared.Result {
	var resumes []*domain.Resume

	if err := r.db.Find(&resumes).Error; err != nil {
		return &shared.Result{Error: err}
	}

	return &shared.Result{Data: resumes}
}

func (r *resumeRepo) FindBySlug(slug string) shared.Result {
	var resume domain.Resume

	if err := r.db.Where(`slug = ?`, slug).Find(&resume).Error; err != nil {
		return shared.Result{Error: err}
	}

	return shared.Result{Data: &resume}
}

func (r *resumeRepo) Save(data *domain.Resume) shared.Result {
	data.Achievements = nil
	data.Experiences = nil
	var resume domain.Resume
	if err := r.db.Where(domain.Resume{Slug: data.Slug}).Assign(data).FirstOrCreate(&resume).Error; err != nil {
		return shared.Result{Error: err}
	}

	return shared.Result{Data: &resume}
}
