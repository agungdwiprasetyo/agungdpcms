package usecase

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
)

// Resume usecase abstraction
type Resume interface {
	FindAll() *shared.Result
	FindBySlug(slug string) shared.Result
	Save(*domain.Resume) shared.Result
}
