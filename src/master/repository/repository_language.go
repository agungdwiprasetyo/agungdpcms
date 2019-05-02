package repository

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/shared/filter"
	"github.com/jinzhu/gorm"
)

type languageGorm struct {
	db *gorm.DB
}

// NewLanguageRepo init master language repo
func NewLanguageRepo(db *gorm.DB) Language {
	return &languageGorm{db}
}

func (r *languageGorm) FindAll(f *filter.Filter) <-chan shared.Result {
	output := make(chan shared.Result)

	go func() {
		defer close(output)
	}()

	return output
}
