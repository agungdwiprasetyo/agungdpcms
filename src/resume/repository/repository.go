package repository

import (
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/agungdwiprasetyo/agungdpcms/shared"
)

// Resume abstraction
type Resume interface {
	FindAll() *shared.Result
	FindBySlug(slug string) *shared.Result
	Save(*domain.Resume) *shared.Result
}

// Achievement abstraction
type Achievement interface {
	FindByResumeID(resumeID int) *shared.Result
}
