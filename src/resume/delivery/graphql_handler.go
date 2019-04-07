package delivery

import (
	"context"
	"fmt"

	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/serializer"
	"github.com/agungdwiprasetyo/agungdpcms/src/shared"
)

// ResumeHandler graphql
type ResumeHandler struct{}

func New() *ResumeHandler {
	return &ResumeHandler{}
}

func (h *ResumeHandler) GetResume(ctx context.Context) (*serializer.ResumeSchema, error) {
	headers := shared.ParseHeaderFromContext(ctx)
	fmt.Println(headers.Get("Authorization"))

	return &serializer.ResumeSchema{Resume: &domain.Resume{ID: 2000, Name: "wkwkwk", Percentage: 323.9}}, nil
}
