package delivery

import (
	"context"

	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/serializer"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/usecase"
	"github.com/agungdwiprasetyo/go-utils/debug"
)

// ResumeHandler graphql
type ResumeHandler struct {
	uc   usecase.Resume
	midd middleware.Middleware
}

// New constructor
func New(uc usecase.Resume, midd middleware.Middleware) *ResumeHandler {
	return &ResumeHandler{uc, midd}
}

// GetAllResume handler
func (h *ResumeHandler) GetAllResume(ctx context.Context) (*serializer.ResumeListSchema, error) {
	h.midd.WithAuth(ctx)

	result := h.uc.FindAll()
	if result.Error != nil {
		return nil, result.Error
	}
	return result.Data.(*serializer.ResumeListSchema), nil
}

// GetResumeBySlug handler
func (h *ResumeHandler) GetResumeBySlug(ctx context.Context, args *ResumeSlugInput) (*serializer.ResumeSchema, error) {
	result := h.uc.FindBySlug(args.Slug)
	if result.Error != nil {
		return nil, result.Error
	}

	return result.Data.(*serializer.ResumeSchema), nil
}

// CreateResume handler
func (h *ResumeHandler) CreateResume(ctx context.Context, args *serializer.ResumeSchema) (*serializer.ResumeSchema, error) {
	h.midd.WithAuth(ctx)

	debug.PrintJSON(args)

	result := h.uc.Save(args.Resume)
	if result.Error != nil {
		return nil, result.Error
	}
	data := result.Data.(*domain.Resume)
	return &serializer.ResumeSchema{Resume: data}, nil
}
