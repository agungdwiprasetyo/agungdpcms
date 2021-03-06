package serializer

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared/meta"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
)

type ResumeSchema struct {
	Resume          *domain.Resume
	ProfileSchema   *ProfileSchema
	AchievementList []*AchievementSchema
	ExperienceList  []*ExperienceSchema
	SkillList       []*SkillSchema
}

func (r *ResumeSchema) ID() int32 {
	return int32(r.Resume.ID)
}

func (r *ResumeSchema) Slug() string {
	return r.Resume.Slug
}

func (r *ResumeSchema) Name() string {
	return r.Resume.Name
}

func (r *ResumeSchema) Profile() *ProfileSchema {
	return r.ProfileSchema
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
