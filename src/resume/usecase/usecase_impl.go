package usecase

import (
	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	rr "github.com/agungdwiprasetyo/agungdpcms/src/resume/repository"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/serializer"
)

type resumeUc struct {
	repo *rr.Repository
}

// NewResumeUsecase constructor
func NewResumeUsecase(conf *config.Config) Resume {
	return &resumeUc{
		repo: rr.NewRepository(conf.DB),
	}
}

func (uc *resumeUc) FindAll() *shared.Result {
	result := uc.repo.Resume.FindAll()
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
	result := uc.repo.Resume.FindBySlug(slug)
	if result.Error != nil {
		return result
	}

	resume := result.Data.(*domain.Resume)

	data := new(serializer.ResumeSchema)
	data.Resume = resume

	achChan := make(chan []*serializer.AchievementSchema)
	expChan := make(chan []*serializer.ExperienceSchema)

	go func() {
		result := uc.repo.Achievement.FindByResumeID(resume.ID)
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
		result = uc.repo.Experience.FindByResumeID(resume.ID)
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

func (uc *resumeUc) Save(data *domain.Resume) (res shared.Result) {
	err := uc.repo.WithTransaction(func(repo *rr.Repository) error {
		achievements, experiences := data.Achievements, data.Experiences

		// save resume data
		res = repo.Resume.Save(data)
		if res.Error != nil {
			return res.Error
		}
		resume := res.Data.(*domain.Resume)

		// save achievement data
		for _, ach := range achievements {
			ach.ResumeID = resume.ID
			res = repo.Achievement.Save(ach)
			if res.Error != nil {
				return res.Error
			}
			resume.Achievements = append(resume.Achievements, res.Data.(*domain.Achievement))
		}

		// save experience data
		for _, exp := range experiences {
			exp.ResumeID = resume.ID
			res = repo.Experience.Save(exp)
			if res.Error != nil {
				return res.Error
			}
			resume.Experiences = append(resume.Experiences, res.Data.(*domain.Experience))
		}

		res.Data = resume
		return nil
	})
	if err != nil {
		res.Error = err
	}

	return res
}
