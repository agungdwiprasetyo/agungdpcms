package usecase

import (
	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	rr "github.com/agungdwiprasetyo/agungdpcms/src/resume/repository"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/serializer"
	"github.com/agungdwiprasetyo/agungdpcms/shared"
)

type resumeUc struct {
	resumeRepo      rr.Resume
	achievementRepo rr.Achievement
}

// NewResumeUsecase constructor
func NewResumeUsecase(conf *config.Config) Resume {
	return &resumeUc{
		resumeRepo:      rr.NewResumeRepository(conf.DB),
		achievementRepo: rr.NewAchievementRepository(conf.DB),
	}
}

func (uc *resumeUc) FindAll() *shared.Result {
	result := uc.resumeRepo.FindAll()
	if result.Error != nil {
		return result
	}

	data := result.Data.([]*domain.Resume)
	fields := make([]*serializer.ResumeSchema, 0)
	for _, d := range data {
		fields = append(fields, &serializer.ResumeSchema{Resume: d})
	}
	return &shared.Result{Data: &serializer.ResumeListSchema{Data: fields}}
}

func (uc *resumeUc) FindBySlug(slug string) *shared.Result {
	result := uc.resumeRepo.FindBySlug(slug)
	if result.Error != nil {
		return result
	}

	resume := result.Data.(*domain.Resume)
	result = uc.achievementRepo.FindByResumeID(resume.ID)
	if result.Error != nil {
		return result
	}

	achievements := result.Data.([]*domain.Achievement)

	data := new(serializer.ResumeSchema)
	data.Resume = resume
	data.AchievementList = make([]*serializer.AchievementSchema, 0)
	for _, ach := range achievements {
		data.AchievementList = append(data.AchievementList, &serializer.AchievementSchema{Achievement: ach})
	}

	return &shared.Result{Data: data}
}

func (uc *resumeUc) Save(data *domain.Resume) *shared.Result {
	return uc.resumeRepo.Save(data)
}
