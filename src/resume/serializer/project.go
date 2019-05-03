package serializer

import "github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"

type ProjectSchema struct {
	Project *domain.Project
}

func (p *ProjectSchema) ID() int32 {
	return int32(p.Project.ID)
}

func (p *ProjectSchema) ResumeID() int32 {
	return int32(p.Project.ResumeID)
}

func (p *ProjectSchema) Name() string {
	return p.Project.Name
}

func (p *ProjectSchema) Description() string {
	return p.Project.Description
}

func (p *ProjectSchema) Date() string {
	return p.Project.Date
}

func (p *ProjectSchema) URL() string {
	return p.Project.URL
}

// func (p *ProjectSchema) Language() string {
// 	return p.Project.Language
// }

func (p *ProjectSchema) Technology() string {
	return p.Project.Technology
}

func (p *ProjectSchema) Repository() string {
	return p.Project.Repository
}
