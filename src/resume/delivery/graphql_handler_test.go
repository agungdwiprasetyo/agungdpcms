package delivery

import (
	"context"
	"fmt"
	"testing"

	middMocks "github.com/agungdwiprasetyo/agungdpcms/middleware/mocks"
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	validatorMocks "github.com/agungdwiprasetyo/agungdpcms/shared/validator/mocks"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/serializer"
	ucMocks "github.com/agungdwiprasetyo/agungdpcms/src/resume/usecase/mocks"
	"github.com/agungdwiprasetyo/go-utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGraphQLHandler_GetAllResume(t *testing.T) {
	type args struct {
		ctx  context.Context
		args *domain.GetAllResumeArgs
	}
	tests := []struct {
		name           string
		args           args
		wantValidation *utils.MultiError
		wantUsecase    shared.Result
		wantErr        bool
	}{
		{
			name:        "Testcase #1: Positive",
			args:        args{ctx: context.Background(), args: &domain.GetAllResumeArgs{}},
			wantUsecase: shared.Result{Data: &serializer.ResumeListSchema{}},
		},
		{
			name:           "Testcase #2: Negative, error validation",
			args:           args{ctx: context.Background(), args: &domain.GetAllResumeArgs{}},
			wantValidation: utils.NewMultiError(),
			wantUsecase:    shared.Result{Data: &serializer.ResumeListSchema{}},
			wantErr:        true,
		},
		{
			name:        "Testcase #3: Negative, error fetch data from usecase",
			args:        args{ctx: context.Background(), args: &domain.GetAllResumeArgs{}},
			wantUsecase: shared.Result{Error: fmt.Errorf("error")},
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resumeUsecase := new(ucMocks.Resume)
			resumeUsecase.On("FindAll", mock.Anything).Return(tt.wantUsecase)

			midd := new(middMocks.Middleware)
			midd.On("WithAuth", mock.Anything).Return(tt.args.ctx)

			v := new(validatorMocks.Validator)
			v.On("Validate", mock.Anything).Return(tt.wantValidation)

			h := NewGraphQLHandler(resumeUsecase, midd, v)
			got, err := h.GetAllResume(tt.args.ctx, tt.args.args)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}

func TestGraphQLHandler_GetResumeBySlug(t *testing.T) {
	type args struct {
		ctx  context.Context
		args *domain.ResumeSlugInput
	}
	tests := []struct {
		name           string
		args           args
		wantValidation *utils.MultiError
		wantUsecase    shared.Result
		wantErr        bool
	}{
		{
			name:        "Testcase #1: Positive",
			args:        args{ctx: context.Background(), args: &domain.ResumeSlugInput{}},
			wantUsecase: shared.Result{Data: &serializer.ResumeSchema{}},
		},
		{
			name:        "Testcase #2: Negative, error fetch data from usecase",
			args:        args{ctx: context.Background(), args: &domain.ResumeSlugInput{}},
			wantUsecase: shared.Result{Error: fmt.Errorf("error")},
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resumeUsecase := new(ucMocks.Resume)
			resumeUsecase.On("FindBySlug", mock.Anything).Return(tt.wantUsecase)

			midd := new(middMocks.Middleware)
			v := new(validatorMocks.Validator)

			h := NewGraphQLHandler(resumeUsecase, midd, v)
			got, err := h.GetResumeBySlug(tt.args.ctx, tt.args.args)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}

func TestGraphQLHandler_CreateResume(t *testing.T) {
	type args struct {
		ctx  context.Context
		args *serializer.ResumeSchema
	}
	tests := []struct {
		name           string
		args           args
		wantValidation *utils.MultiError
		wantUsecase    shared.Result
		wantErr        bool
	}{
		{
			name:        "Testcase #1: Positive",
			args:        args{ctx: context.Background(), args: &serializer.ResumeSchema{}},
			wantUsecase: shared.Result{Data: &domain.Resume{}},
		},
		{
			name:           "Testcase #2: Negative, error validation",
			args:           args{ctx: context.Background(), args: &serializer.ResumeSchema{}},
			wantValidation: utils.NewMultiError(),
			wantUsecase:    shared.Result{Data: &domain.Resume{}},
			wantErr:        true,
		},
		{
			name:        "Testcase #3: Negative, error fetch data from usecase",
			args:        args{ctx: context.Background(), args: &serializer.ResumeSchema{}},
			wantUsecase: shared.Result{Error: fmt.Errorf("error")},
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resumeUsecase := new(ucMocks.Resume)
			resumeUsecase.On("Save", mock.Anything).Return(tt.wantUsecase)

			midd := new(middMocks.Middleware)
			midd.On("WithAuth", mock.Anything).Return(tt.args.ctx)

			v := new(validatorMocks.Validator)
			v.On("Validate", mock.Anything).Return(tt.wantValidation)

			h := NewGraphQLHandler(resumeUsecase, midd, v)
			got, err := h.CreateResume(tt.args.ctx, tt.args.args)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}

func TestGraphQLHandler_Remove(t *testing.T) {
	generateInt := func(n int) *int32 {
		t := int32(n)
		return &t
	}
	type args struct {
		ctx  context.Context
		args *domain.RemoveArgs
	}
	tests := []struct {
		name                   string
		args                   args
		wantUsecaseAchievement shared.Result
		wantUsecaseExperience  shared.Result
		wantUsecaseSkill       shared.Result
		wantErr                bool
	}{
		{
			name: "Testcase #1: Positive",
			args: args{ctx: context.Background(), args: &domain.RemoveArgs{}},
		},
		{
			name: "Testcase #2: Negative, error validation",
			args: args{ctx: context.Background(), args: &domain.RemoveArgs{
				AchievementID: generateInt(10), ExperienceID: generateInt(10), SkillID: generateInt(10),
			}},
			wantUsecaseAchievement: shared.Result{Error: fmt.Errorf("error")},
			wantUsecaseExperience:  shared.Result{Error: fmt.Errorf("error")},
			wantUsecaseSkill:       shared.Result{Error: fmt.Errorf("error")},
			wantErr:                true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resumeUsecase := new(ucMocks.Resume)
			resumeUsecase.On("RemoveAchievement", mock.Anything).Return(tt.wantUsecaseAchievement)
			resumeUsecase.On("RemoveExperience", mock.Anything).Return(tt.wantUsecaseExperience)
			resumeUsecase.On("RemoveSkill", mock.Anything).Return(tt.wantUsecaseSkill)

			midd := new(middMocks.Middleware)
			midd.On("WithAuth", mock.Anything).Return(tt.args.ctx)

			v := new(validatorMocks.Validator)

			h := NewGraphQLHandler(resumeUsecase, midd, v)
			got, err := h.Remove(tt.args.ctx, tt.args.args)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}
