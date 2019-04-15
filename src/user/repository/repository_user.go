package repository

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/user/domain"
	"github.com/jinzhu/gorm"
)

type userGorm struct {
	db *gorm.DB
}

// NewUserRepository construct user repo
func NewUserRepository(db *gorm.DB) User {
	return &userGorm{db}
}

func (r *userGorm) FindByUsername(username string) <-chan shared.Result {
	output := make(chan shared.Result)

	go func() {
		defer close(output)

		var user domain.User
		if err := r.db.Where(&domain.User{Username: username}).First(&user).Error; err != nil {
			output <- shared.Result{Error: err}
			return
		}

		output <- shared.Result{Data: &user}
	}()

	return output
}
