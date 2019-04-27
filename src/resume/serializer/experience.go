package serializer

import (
	"time"

	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
)

type ExperienceSchema struct {
	Experience *domain.Experience
}

func (e *ExperienceSchema) ID() int32 {
	return int32(e.Experience.ID)
}

func (e *ExperienceSchema) ResumeID() int32 {
	return int32(e.Experience.ResumeID)
}

func (e *ExperienceSchema) Title() string {
	return e.Experience.Title
}

func (e *ExperienceSchema) Company() string {
	return e.Experience.Company
}

func (e *ExperienceSchema) Location() string {
	return e.Experience.Location
}

func (e *ExperienceSchema) StartDate() string {
	return e.Experience.StartDate.Format(time.RFC3339)
}

func (e *ExperienceSchema) EndDate() string {
	return e.Experience.EndDate.Format(time.RFC3339)
}
