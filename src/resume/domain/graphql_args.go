package domain

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared/filter"
)

// GetAllResumeArgs args
type GetAllResumeArgs struct {
	Filter filter.Filter
}

// ResumeSlugInput args
type ResumeSlugInput struct {
	Slug string
}

// RemoveArgs args
type RemoveArgs struct {
	AchievementID, ExperienceID, SkillID *int32
}
