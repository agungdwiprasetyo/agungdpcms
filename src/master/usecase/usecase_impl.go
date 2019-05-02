package usecase

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/shared/filter"
	mr "github.com/agungdwiprasetyo/agungdpcms/src/master/repository"
)

type masterUc struct {
	repo *mr.Repository
}

// NewMasterUsecase construct master usecase
func NewMasterUsecase(repo *mr.Repository) Master {
	return &masterUc{
		repo: repo,
	}
}

func (uc *masterUc) FindAllLanguage(f *filter.Filter) (res shared.Result) {
	return
}
