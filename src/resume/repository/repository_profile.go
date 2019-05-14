package repository

import (
	"fmt"

	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/jinzhu/gorm"
)

type profileRepo struct {
	db *gorm.DB
}

// NewProfileRepository construct new profile repo
func NewProfileRepository(db *gorm.DB) Profile {
	return &profileRepo{db}
}

func (r *profileRepo) FindByResumeID(resumeID int) <-chan *domain.Profile {
	output := make(chan *domain.Profile)

	go func() {
		defer close(output)

		var profile domain.Profile
		if err := r.db.Where(domain.Profile{ResumeID: resumeID}).Find(&profile).Error; err != nil {
			panic(err)
		}
		output <- &profile
	}()

	return output
}

func (r *profileRepo) Save(data *domain.Profile) (res shared.Result) {
	var profile domain.Profile
	where := domain.Profile{
		ResumeID: data.ResumeID, Fullname: data.Fullname,
	}
	if err := r.db.Where(where).Assign(data).FirstOrCreate(&profile).Error; err != nil {
		res.Error = err
	}
	res.Data = &profile
	return
}

func (r *profileRepo) Remove(data *domain.Profile) (res shared.Result) {
	db := r.db.Delete(data)
	if err := db.Error; err != nil {
		res.Error = err
	}
	if affected := db.RowsAffected; affected == 0 {
		res.Error = fmt.Errorf("data with id=%d not found", data.ID)
	}
	return
}
