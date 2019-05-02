package usecase

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/shared/filter"
)

// Master usecase abstraction
type Master interface {
	FindAllLanguage(*filter.Filter) shared.Result
}
