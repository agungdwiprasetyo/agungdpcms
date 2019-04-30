package serializer

import "github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"

type SkillSchema struct {
	Skill *domain.Skill
}

func (s *SkillSchema) ID() int32 {
	return int32(s.Skill.ID)
}

func (s *SkillSchema) ResumeID() int32 {
	return int32(s.Skill.ResumeID)
}

func (s *SkillSchema) Type() string {
	return s.Skill.Type
}

func (s *SkillSchema) Name() string {
	return s.Skill.Name
}

func (s *SkillSchema) Description() string {
	return s.Skill.Description
}

func (s *SkillSchema) Percentage() float64 {
	return s.Skill.Percentage
}
