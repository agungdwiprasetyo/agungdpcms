package serializer

import "github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"

type ProfileSchema struct {
	Profile *domain.Profile
}

func (p *ProfileSchema) ID() int32 {
	return int32(p.Profile.ID)
}

func (p *ProfileSchema) ResumeID() int32 {
	return int32(p.Profile.ResumeID)
}

func (p *ProfileSchema) Fullname() string {
	return p.Profile.Fullname
}

func (p *ProfileSchema) Religion() string {
	return p.Profile.Religion
}

func (p *ProfileSchema) Hobby() string {
	return p.Profile.Hobby
}

func (p *ProfileSchema) Github() string {
	return p.Profile.Github
}

func (p *ProfileSchema) Linkedin() string {
	return p.Profile.Linkedin
}

func (p *ProfileSchema) Instagram() string {
	return p.Profile.Instagram
}

func (p *ProfileSchema) Facebook() string {
	return p.Profile.Facebook
}
