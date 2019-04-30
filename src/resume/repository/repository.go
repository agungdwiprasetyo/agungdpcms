package repository

import (
	"fmt"

	"github.com/agungdwiprasetyo/agungdpcms/shared/filter"

	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/jinzhu/gorm"
)

// declare all repository
type (
	// Resume abstraction
	Resume interface {
		FindAll(*filter.Filter) shared.Result
		Count(*domain.Resume) int
		FindBySlug(slug string) shared.Result
		Save(*domain.Resume) shared.Result
	}

	// Profile abstraction
	Profile interface {
		FindByResumeID(resumeID int) shared.Result
		Save(data *domain.Profile) shared.Result
		Remove(data *domain.Profile) shared.Result
	}

	// Achievement abstraction
	Achievement interface {
		FindByResumeID(resumeID int) shared.Result
		Save(data *domain.Achievement) shared.Result
		Remove(data *domain.Achievement) shared.Result
	}

	// Experience abstraction
	Experience interface {
		FindByResumeID(resumeID int) shared.Result
		Save(data *domain.Experience) shared.Result
		Remove(data *domain.Experience) shared.Result
	}

	// Skill abstraction
	Skill interface {
		FindByResumeID(resumeID int) shared.Result
		Save(data *domain.Skill) shared.Result
		Remove(data *domain.Skill) shared.Result
	}
)

// Repository parent
type Repository struct {
	db          *gorm.DB
	Resume      Resume
	Profile     Profile
	Achievement Achievement
	Experience  Experience
	Skill       Skill
}

// NewRepository repository constructor
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db:          db,
		Resume:      NewResumeRepository(db),
		Profile:     NewProfileRepository(db),
		Achievement: NewAchievementRepository(db),
		Experience:  NewExperienceRepository(db),
		Skill:       NewSkillRepository(db),
	}
}

// WithTransaction run transaction for each repository
func (r *Repository) WithTransaction(txFunc func(*Repository) error) (err error) {
	db := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
			db.Rollback()
		} else if err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}()

	// reinit new repository in different memory address
	manager := NewRepository(db)
	err = txFunc(manager)
	return
}
