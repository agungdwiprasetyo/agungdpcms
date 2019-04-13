package serializer

import (
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	graphql "github.com/graph-gophers/graphql-go"
)

type AchievementSchema struct {
	Achievement *domain.Achievement
}

func (a *AchievementSchema) ID() graphql.ID {
	return graphql.ID(a.Achievement.ID)
}

func (a *AchievementSchema) ResumeID() graphql.ID {
	return graphql.ID(a.Achievement.ResumeID)
}

func (a *AchievementSchema) Name() string {
	return a.Achievement.Name
}

func (a *AchievementSchema) Appreciator() string {
	return a.Achievement.Appreciator
}

func (a *AchievementSchema) Year() int32 {
	return int32(a.Achievement.Year)
}

func (a *AchievementSchema) Picture() string {
	return a.Achievement.Picture
}

type AchievementListSchema struct {
	Data []*AchievementSchema
}

func (a *AchievementListSchema) Achievements() []*AchievementSchema {
	return a.Data
}
