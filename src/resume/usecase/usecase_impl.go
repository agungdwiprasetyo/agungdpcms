package usecase

import (
	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	rr "github.com/agungdwiprasetyo/agungdpcms/src/resume/repository"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/serializer"
	"github.com/agungdwiprasetyo/go-utils/debug"
)

type resumeUc struct {
	resumeRepo      rr.Resume
	achievementRepo rr.Achievement
	experienceRepo  rr.Experience
}

// NewResumeUsecase constructor
func NewResumeUsecase(conf *config.Config) Resume {
	return &resumeUc{
		resumeRepo:      rr.NewResumeRepository(conf.DB),
		achievementRepo: rr.NewAchievementRepository(conf.DB),
		experienceRepo:  rr.NewExperienceRepository(conf.DB),
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

func (uc *resumeUc) FindBySlug(slug string) shared.Result {
	result := uc.resumeRepo.FindBySlug(slug)
	if result.Error != nil {
		return result
	}

	resume := result.Data.(*domain.Resume)

	data := new(serializer.ResumeSchema)
	data.Resume = resume

	achChan := make(chan []*serializer.AchievementSchema)
	expChan := make(chan []*serializer.ExperienceSchema)

	go func() {
		result := uc.achievementRepo.FindByResumeID(resume.ID)
		if result.Error != nil {
			return
		}

		var achievements []*serializer.AchievementSchema
		for _, ach := range result.Data.([]*domain.Achievement) {
			achievements = append(achievements, &serializer.AchievementSchema{Achievement: ach})
		}
		achChan <- achievements
	}()

	go func() {
		result = uc.experienceRepo.FindByResumeID(resume.ID)
		if result.Error != nil {
			return
		}

		var experiences []*serializer.ExperienceSchema
		for _, exp := range result.Data.([]*domain.Experience) {
			experiences = append(experiences, &serializer.ExperienceSchema{Experience: exp})
		}
		expChan <- experiences
	}()

	data.AchievementList = <-achChan
	data.ExperienceList = <-expChan

	return shared.Result{Data: data}
}

func (uc *resumeUc) Save(data *domain.Resume) *shared.Result {
	debug.PrintJSON(data)
	res := uc.resumeRepo.Save(data)
	if res.Error != nil {
		debug.Println(res.Error)
	}
	return res
}
