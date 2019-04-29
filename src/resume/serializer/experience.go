package serializer

import (
	"time"

	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	graphql "github.com/graph-gophers/graphql-go"
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

func (e *ExperienceSchema) StartDate() graphql.Time {
	t, _ := time.Parse(time.RFC3339, e.Experience.StartDate)
	return graphql.Time{Time: t}
}

func (e *ExperienceSchema) EndDate() graphql.Time {
	t, _ := time.Parse(time.RFC3339, e.Experience.EndDate)
	return graphql.Time{Time: t}
}
