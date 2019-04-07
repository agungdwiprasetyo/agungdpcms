package serializer

import (
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	graphql "github.com/graph-gophers/graphql-go"
)

type ResumeSchema struct {
	Resume *domain.Resume
}

func (r *ResumeSchema) ID() graphql.ID {
	return graphql.ID(r.Resume.ID)
}

func (r *ResumeSchema) Name() string {
	return r.Resume.Name
}

func (r *ResumeSchema) Percentage() float64 {
	return r.Resume.Percentage
}
