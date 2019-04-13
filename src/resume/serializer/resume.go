package serializer

import (
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	graphql "github.com/graph-gophers/graphql-go"
)

type ResumeSchema struct {
	Resume          *domain.Resume
	AchievementList []*AchievementSchema
}

func (r *ResumeSchema) ID() graphql.ID {
	return graphql.ID(r.Resume.ID)
}

func (r *ResumeSchema) Slug() string {
	return r.Resume.Slug
}

func (r *ResumeSchema) Name() string {
	return r.Resume.Name
}

func (r *ResumeSchema) Achievements() []*AchievementSchema {
	return r.AchievementList
}

type ResumeListSchema struct {
	Data []*ResumeSchema
}

// Resumes method
func (r *ResumeListSchema) Resumes() []*ResumeSchema {
	return r.Data
}
