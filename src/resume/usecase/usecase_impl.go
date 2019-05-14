package usecase

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/shared/filter"
	"github.com/agungdwiprasetyo/agungdpcms/shared/meta"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/repository"
	rr "github.com/agungdwiprasetyo/agungdpcms/src/resume/repository"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/serializer"
)

type resumeUc struct {
	repo *rr.Repository
}

// NewResumeUsecase constructor
func NewResumeUsecase(repo *repository.Repository) Resume {
	return &resumeUc{
		repo: repo,
	}
}

func (uc *resumeUc) FindAll(filter *filter.Filter) shared.Result {
	filter.CalculateOffset()
	result := uc.repo.Resume.FindAll(filter)
	if result.Error != nil {
		return result
	}

	data := result.Data.([]*domain.Resume)
	fields := make([]*serializer.ResumeSchema, 0)
	for _, d := range data {
		fields = append(fields, &serializer.ResumeSchema{Resume: d})
	}

	count := uc.repo.Resume.Count(&domain.Resume{})

	m := &meta.Meta{Page: int(filter.Page), Limit: int(filter.Limit), TotalRecords: count}
	m.CalculatePages()

	return shared.Result{Data: &serializer.ResumeListSchema{
		M: &meta.MetaSchema{Meta: m}, Data: fields,
	}}
}

func (uc *resumeUc) FindBySlug(slug string) shared.Result {
	result := uc.repo.Resume.FindBySlug(slug)
	if result.Error != nil {
		return result
	}

	resume := result.Data.(*domain.Resume)

	data := new(serializer.ResumeSchema)
	data.Resume = resume

	profileChan := uc.repo.Profile.FindByResumeID(resume.ID)
	achChan := uc.repo.Achievement.FindByResumeID(resume.ID)
	expChan := uc.repo.Experience.FindByResumeID(resume.ID)
	skillChan := uc.repo.Skill.FindByResumeID(resume.ID)

	data.ProfileSchema = &serializer.ProfileSchema{Profile: <-profileChan}
	for _, ach := range <-achChan {
		data.AchievementList = append(data.AchievementList, &serializer.AchievementSchema{Achievement: ach})
	}
	for _, exp := range <-expChan {
		data.ExperienceList = append(data.ExperienceList, &serializer.ExperienceSchema{Experience: exp})
	}
	for _, skill := range <-skillChan {
		data.SkillList = append(data.SkillList, &serializer.SkillSchema{Skill: skill})
	}

	return shared.Result{Data: data}
}

func (uc *resumeUc) Save(data *domain.Resume) (res shared.Result) {
	err := uc.repo.WithTransaction(func(repo *rr.Repository) error {
		achievements, experiences, skills := data.Achievements, data.Experiences, data.Skills
		profile := data.Profile
		data.EmptyChild()

		// save resume data
		res = repo.Resume.Save(data)
		if res.Error != nil {
			return res.Error
		}
		resume := res.Data.(*domain.Resume)

		if profile != nil {
			profile.ResumeID = resume.ID
			res = uc.repo.Profile.Save(profile)
			if res.Error != nil {
				return res.Error
			}
		}

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

		// save skills data
		for _, skill := range skills {
			skill.ResumeID = resume.ID
			res = repo.Skill.Save(skill)
			if res.Error != nil {
				return res.Error
			}
			resume.Skills = append(resume.Skills, res.Data.(*domain.Skill))
		}

		res.Data = resume
		return nil
	})
	if err != nil {
		res.Error = err
	}

	return res
}

func (uc *resumeUc) RemoveAchievement(id int) (res shared.Result) {
	ach := domain.Achievement{ID: id}
	return uc.repo.Achievement.Remove(&ach)
}

func (uc *resumeUc) RemoveExperience(id int) (res shared.Result) {
	exp := domain.Experience{ID: id}
	return uc.repo.Experience.Remove(&exp)
}

func (uc *resumeUc) RemoveSkill(id int) (res shared.Result) {
	skill := domain.Skill{ID: id}
	return uc.repo.Skill.Remove(&skill)
}
