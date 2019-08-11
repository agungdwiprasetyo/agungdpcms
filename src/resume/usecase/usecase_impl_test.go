package usecase

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/shared/filter"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/repository"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func generateProfile(data *domain.Profile) <-chan *domain.Profile {
	output := make(chan *domain.Profile)
	go func() {
		defer close(output)
		output <- data
	}()
	return output
}

func generateAchievement(data []*domain.Achievement) <-chan []*domain.Achievement {
	output := make(chan []*domain.Achievement)
	go func() {
		defer close(output)
		output <- data
	}()
	return output
}

func generateExperience(data []*domain.Experience) <-chan []*domain.Experience {
	output := make(chan []*domain.Experience)
	go func() {
		defer close(output)
		output <- data
	}()
	return output
}

func generateSkill(data []*domain.Skill) <-chan []*domain.Skill {
	output := make(chan []*domain.Skill)
	go func() {
		defer close(output)
		output <- data
	}()
	return output
}

func Test_resumeUc_FindAll(t *testing.T) {
	type repoResult struct {
		resume shared.Result
	}
	tests := []struct {
		name           string
		filter         *filter.Filter
		wantRepoResult repoResult
		wantError      bool
	}{
		{
			name:   "Testcase #1: Positive",
			filter: &filter.Filter{},
			wantRepoResult: repoResult{
				resume: shared.Result{Data: []*domain.Resume{&domain.Resume{}}},
			},
			wantError: false,
		},
		{
			name:   "Testcase #2: Negative",
			filter: &filter.Filter{},
			wantRepoResult: repoResult{
				resume: shared.Result{Error: fmt.Errorf("error")},
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resumeRepoMock := new(mocks.Resume)
			resumeRepoMock.On("FindAll", mock.Anything).Return(tt.wantRepoResult.resume)
			resumeRepoMock.On("Count", mock.Anything).Return(10)

			uc := NewResumeUsecase(&repository.Repository{
				Resume: resumeRepoMock,
			})

			got := uc.FindAll(tt.filter)
			if tt.wantError {
				assert.Error(t, got.Error)
			} else {
				assert.NoError(t, got.Error)
			}
		})
	}
}

func Test_resumeUc_FindBySlug(t *testing.T) {
	type repoResult struct {
		resume shared.Result
	}
	tests := []struct {
		name           string
		wantRepoResult repoResult
		wantError      bool
	}{
		{
			name: "Testcase #1: Positive",
			wantRepoResult: repoResult{
				resume: shared.Result{Data: &domain.Resume{}},
			},
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			wantRepoResult: repoResult{
				resume: shared.Result{Error: fmt.Errorf("error")},
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resumeRepoMock := new(mocks.Resume)
			resumeRepoMock.On("FindBySlug", mock.Anything).Return(tt.wantRepoResult.resume)

			profileRepoMock := new(mocks.Profile)
			profileRepoMock.On("FindByResumeID", mock.Anything).Return(generateProfile(&domain.Profile{}))
			achRepoMock := new(mocks.Achievement)
			achRepoMock.On("FindByResumeID", mock.Anything).Return(generateAchievement([]*domain.Achievement{&domain.Achievement{}}))
			expRepoMock := new(mocks.Experience)
			expRepoMock.On("FindByResumeID", mock.Anything).Return(generateExperience([]*domain.Experience{&domain.Experience{}}))
			skillRepoMock := new(mocks.Skill)
			skillRepoMock.On("FindByResumeID", mock.Anything).Return(generateSkill([]*domain.Skill{&domain.Skill{}}))

			uc := NewResumeUsecase(&repository.Repository{
				Resume: resumeRepoMock, Profile: profileRepoMock, Achievement: achRepoMock, Experience: expRepoMock, Skill: skillRepoMock,
			})

			got := uc.FindBySlug("agungdp")
			if tt.wantError {
				assert.Error(t, got.Error)
			} else {
				assert.NoError(t, got.Error)
			}
		})
	}
}

func Test_resumeUc_Save(t *testing.T) {
	type repoResult struct {
		resume, profile, achievement, experience, skill shared.Result
	}
	tests := []struct {
		name           string
		wantRepoResult repoResult
		wantError      bool
	}{
		{
			name: "Testcase #1. Positive",
			wantRepoResult: repoResult{
				resume:      shared.Result{Data: &domain.Resume{}},
				profile:     shared.Result{Data: &domain.Profile{}},
				achievement: shared.Result{Data: &domain.Achievement{}},
				experience:  shared.Result{Data: &domain.Experience{}},
				skill:       shared.Result{Data: &domain.Skill{}},
			},
		},
		{
			name: "Testcase #2. Negative, error get resume data",
			wantRepoResult: repoResult{
				resume: shared.Result{Error: errors.New("err")},
			},
			wantError: true,
		},
		{
			name: "Testcase #3. Negative, error get profile data",
			wantRepoResult: repoResult{
				resume:  shared.Result{Data: &domain.Resume{}},
				profile: shared.Result{Error: errors.New("err")},
			},
			wantError: true,
		},
		{
			name: "Testcase #4. Negative, error get achievement data",
			wantRepoResult: repoResult{
				resume:      shared.Result{Data: &domain.Resume{}},
				profile:     shared.Result{Data: &domain.Profile{}},
				achievement: shared.Result{Error: errors.New("err")},
			},
			wantError: true,
		},
		{
			name: "Testcase #5. Negative, error get experience data",
			wantRepoResult: repoResult{
				resume:      shared.Result{Data: &domain.Resume{}},
				profile:     shared.Result{Data: &domain.Profile{}},
				achievement: shared.Result{Data: &domain.Achievement{}},
				experience:  shared.Result{Error: errors.New("err")},
			},
			wantError: true,
		},
		{
			name: "Testcase #6. Negative, error get skill data",
			wantRepoResult: repoResult{
				resume:      shared.Result{Data: &domain.Resume{}},
				profile:     shared.Result{Data: &domain.Profile{}},
				achievement: shared.Result{Data: &domain.Achievement{}},
				experience:  shared.Result{Data: &domain.Experience{}},
				skill:       shared.Result{Error: errors.New("err")},
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resumeRepoMock := new(mocks.Resume)
			resumeRepoMock.On("Save", mock.Anything).Return(tt.wantRepoResult.resume)
			profileRepoMock := new(mocks.Profile)
			profileRepoMock.On("Save", mock.Anything).Return(tt.wantRepoResult.profile)
			achRepoMock := new(mocks.Achievement)
			achRepoMock.On("Save", mock.Anything).Return(tt.wantRepoResult.achievement)
			expRepoMock := new(mocks.Experience)
			expRepoMock.On("Save", mock.Anything).Return(tt.wantRepoResult.experience)
			skillRepoMock := new(mocks.Skill)
			skillRepoMock.On("Save", mock.Anything).Return(tt.wantRepoResult.skill)

			repo := &repository.Repository{
				Resume:      resumeRepoMock,
				Profile:     profileRepoMock,
				Achievement: achRepoMock,
				Experience:  expRepoMock,
				Skill:       skillRepoMock,
			}

			var guard *monkey.PatchGuard
			guard = monkey.PatchInstanceMethod(reflect.TypeOf(repo), "WithTransaction", func(_ *repository.Repository, txFunc func(*repository.Repository) error) error {
				guard.Unpatch()
				defer guard.Restore()
				return txFunc(repo)
			})

			uc := NewResumeUsecase(repo)
			got := uc.Save(&domain.Resume{
				Profile: &domain.Profile{}, Achievements: []*domain.Achievement{&domain.Achievement{}}, Skills: []*domain.Skill{&domain.Skill{}},
				Experiences: []*domain.Experience{&domain.Experience{}},
			})
			if tt.wantError {
				assert.Error(t, got.Error)
			} else {
				assert.NoError(t, got.Error)
			}
		})
	}
}

func Test_resumeUc_Remove(t *testing.T) {
	achRepoMock := new(mocks.Achievement)
	achRepoMock.On("Remove", mock.Anything).Return(shared.Result{})
	expRepoMock := new(mocks.Experience)
	expRepoMock.On("Remove", mock.Anything).Return(shared.Result{})
	skillRepoMock := new(mocks.Skill)
	skillRepoMock.On("Remove", mock.Anything).Return(shared.Result{})

	uc := NewResumeUsecase(&repository.Repository{
		Achievement: achRepoMock, Experience: expRepoMock, Skill: skillRepoMock,
	})
	t.Run("Test Remove Achievement", func(t *testing.T) {
		got := uc.RemoveAchievement(1)
		assert.NoError(t, got.Error)
	})
	t.Run("Test Remove Experience", func(t *testing.T) {
		got := uc.RemoveExperience(1)
		assert.NoError(t, got.Error)
	})
	t.Run("Test Remove Skill", func(t *testing.T) {
		got := uc.RemoveSkill(1)
		assert.NoError(t, got.Error)
	})
}
