package repository

import (
	"fmt"

	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/shared/filter"
	"github.com/agungdwiprasetyo/agungdpcms/src/master/domain"
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

		var languages []*domain.Language
		if err := r.db.Limit(f.Limit).Offset(f.Offset).Order(f.SortBy + " " + f.Sort).Find(&languages).Error; err != nil {
			output <- shared.Result{Error: err}
			return
		}

		output <- shared.Result{Data: languages}
	}()

	return output
}

func (r *languageGorm) FindByType(t string) <-chan shared.Result {
	output := make(chan shared.Result)

	go func() {
		defer close(output)

		var languages []*domain.Language
		if err := r.db.Where(domain.Language{Type: t}).Find(&languages).Error; err != nil {
			output <- shared.Result{Error: err}
			return
		}

		output <- shared.Result{Data: languages}
	}()

	return output
}

func (r *languageGorm) FindByID(id int) <-chan shared.Result {
	output := make(chan shared.Result)

	go func() {
		defer close(output)

		var language domain.Language
		if err := r.db.Where(domain.Language{ID: id}).First(&language).Error; err != nil {
			output <- shared.Result{Error: err}
			return
		}

		output <- shared.Result{Data: language}
	}()

	return output
}

func (r *languageGorm) Save(data *domain.Language) <-chan shared.Result {
	output := make(chan shared.Result)

	go func() {
		defer close(output)

		var lang domain.Language
		where := domain.Language{
			Type: data.Type, Name: data.Name,
		}
		if err := r.db.Where(where).Assign(data).FirstOrCreate(&lang).Error; err != nil {
			output <- shared.Result{Error: err}
			return
		}

		output <- shared.Result{Data: &lang}
	}()

	return output
}

func (r *languageGorm) Remove(data *domain.Language) <-chan shared.Result {
	output := make(chan shared.Result)

	go func() {
		defer close(output)

		var res shared.Result
		db := r.db.Delete(data)
		if err := db.Error; err != nil {
			res.Error = err
		}
		if affected := db.RowsAffected; affected == 0 {
			res.Error = fmt.Errorf("data with id=%d not found", data.ID)
		}

		res.Data = data
		output <- res
	}()

	return output
}
