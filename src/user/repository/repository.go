package repository

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/jinzhu/gorm"
)

// User abstraction
type (
	User interface {
		FindByUsername(username string) <-chan shared.Result
	}

	// Role abstraction
	Role interface {
		FindByID(id int) <-chan shared.Result
	}
)

// Repository parent all user repo
type Repository struct {
	User User
	Role Role
}

// NewRepository init new user repo
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
		Role: NewRoleRepository(db),
	}
}
