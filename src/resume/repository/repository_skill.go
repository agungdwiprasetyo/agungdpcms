package repository

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/jinzhu/gorm"
)

type skillRepo struct {
	db *gorm.DB
}

// NewSkillRepository construct new resume repo
func NewSkillRepository(db *gorm.DB) Skill {
	return &skillRepo{db}
}

func (r *skillRepo) FindByResumeID(resumeID int) (res shared.Result) {
	var skills []*domain.Skill

	if err := r.db.Where(domain.Skill{ResumeID: resumeID}).Find(&skills).Error; err != nil {
		res.Error = err
		return
	}

	res.Data = skills
	return
}

func (r *skillRepo) Save(data *domain.Skill) shared.Result {
	var skill domain.Skill
	where := domain.Skill{
		Type: data.Type, Name: data.Name,
	}
	if err := r.db.Where(where).Assign(data).FirstOrCreate(&skill).Error; err != nil {
		return shared.Result{Error: err}
	}

	return shared.Result{Data: &skill}
}
