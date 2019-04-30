package serializer

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared/meta"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	graphql "github.com/graph-gophers/graphql-go"
)

type ResumeSchema struct {
	Resume          *domain.Resume
	AchievementList []*AchievementSchema
	ExperienceList  []*ExperienceSchema
	SkillList       []*SkillSchema
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

func (r *ResumeSchema) Experiences() []*ExperienceSchema {
	return r.ExperienceList
}

func (r *ResumeSchema) Skills() []*SkillSchema {
	return r.SkillList
}

type ResumeListSchema struct {
	M    *meta.MetaSchema
	Data []*ResumeSchema
}

func (r *ResumeListSchema) Meta() *meta.MetaSchema {
	return r.M
}

// Resumes method
func (r *ResumeListSchema) Resumes() []*ResumeSchema {
	return r.Data
}
