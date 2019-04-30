package delivery

import (
	"context"

	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/serializer"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/usecase"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/validation"
	"github.com/agungdwiprasetyo/go-utils"
)

// ResumeHandler graphql
type ResumeHandler struct {
	uc        usecase.Resume
	midd      middleware.Middleware
	validator *validation.Validator
}

// New constructor
func New(uc usecase.Resume, midd middleware.Middleware) *ResumeHandler {
	return &ResumeHandler{
		uc:        uc,
		midd:      midd,
		validator: validation.New(),
	}
}

// GetAllResume handler
func (h *ResumeHandler) GetAllResume(ctx context.Context, args *domain.GetAllResumeArgs) (*serializer.ResumeListSchema, error) {
	h.midd.WithAuth(ctx)
	if err := h.validator.Validate(args.Filter); err != nil {
		return nil, err
	}

	result := h.uc.FindAll(&args.Filter)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.Data.(*serializer.ResumeListSchema), nil
}

// GetResumeBySlug handler
func (h *ResumeHandler) GetResumeBySlug(ctx context.Context, args *domain.ResumeSlugInput) (*serializer.ResumeSchema, error) {
	result := h.uc.FindBySlug(args.Slug)
	if result.Error != nil {
		return nil, result.Error
	}

	return result.Data.(*serializer.ResumeSchema), nil
}

// CreateResume handler
func (h *ResumeHandler) CreateResume(ctx context.Context, args *serializer.ResumeSchema) (*serializer.ResumeSchema, error) {
	h.midd.WithAuth(ctx)
	if err := h.validator.Validate(args.Resume); err != nil {
		return nil, err
	}

	result := h.uc.Save(args.Resume)
	if result.Error != nil {
		return nil, result.Error
	}
	data := result.Data.(*domain.Resume)
	return &serializer.ResumeSchema{Resume: data}, nil
}

// Remove handler
func (h *ResumeHandler) Remove(ctx context.Context, args *domain.RemoveArgs) (string, error) {
	h.midd.WithAuth(ctx)

	multiError := utils.NewMultiError()
	if args.AchievementID != nil {
		res := h.uc.RemoveAchievement(int(*args.AchievementID))
		if res.Error != nil {
			multiError.Append("achievement", res.Error)
		}
	}
	if args.ExperienceID != nil {
		res := h.uc.RemoveExperience(int(*args.ExperienceID))
		if res.Error != nil {
			multiError.Append("experience", res.Error)
		}
	}
	if args.SkillID != nil {
		res := h.uc.RemoveSkill(int(*args.SkillID))
		if res.Error != nil {
			multiError.Append("skill", res.Error)
		}
	}

	if multiError.HasError() {
		return "", multiError
	}

	return "Success", nil
}
