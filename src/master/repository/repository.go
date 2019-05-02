package repository

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/shared/filter"
	"github.com/jinzhu/gorm"
)

type (
	// Language repo abstraction
	Language interface {
		FindAll(*filter.Filter) <-chan shared.Result
	}
)

// Repository parent all master repo
type Repository struct {
	Language Language
}

// NewRepository init master repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Language: NewLanguageRepo(db),
	}
}
