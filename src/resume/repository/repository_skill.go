package repository

import (
	"fmt"

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

func (r *skillRepo) FindByResumeID(resumeID int) <-chan []*domain.Skill {
	output := make(chan []*domain.Skill)

	go func() {
		defer close(output)

		var skills []*domain.Skill
		if err := r.db.Where(domain.Skill{ResumeID: resumeID}).Find(&skills).Error; err != nil {
			panic(err)
		}
		output <- skills
	}()

	return output
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

func (r *skillRepo) Remove(data *domain.Skill) (res shared.Result) {
	db := r.db.Delete(data)
	if err := db.Error; err != nil {
		res.Error = err
	}
	if affected := db.RowsAffected; affected == 0 {
		res.Error = fmt.Errorf("data with id=%d not found", data.ID)
	}
	return
}
