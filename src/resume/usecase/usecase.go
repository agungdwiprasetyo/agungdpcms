package usecase

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/shared/filter"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
)

// Resume usecase abstraction
type Resume interface {
	FindAll(*filter.Filter) shared.Result
	FindBySlug(slug string) shared.Result
	Save(*domain.Resume) shared.Result

	RemoveAchievement(id int) shared.Result
	RemoveExperience(id int) shared.Result
	RemoveSkill(id int) shared.Result
}
