package repository

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/src/user/domain"
	"github.com/jinzhu/gorm"
)

type roleGorm struct {
	db *gorm.DB
}

// NewRoleRepository construct role repo
func NewRoleRepository(db *gorm.DB) Role {
	db.AutoMigrate(&domain.Role{})
	return &roleGorm{db}
}

func (r *roleGorm) FindByID(id int) <-chan shared.Result {
	output := make(chan shared.Result)

	go func() {
		defer close(output)

		var role domain.Role
		if err := r.db.First(&role, id).Error; err != nil {
			output <- shared.Result{Error: err}
			return
		}

		output <- shared.Result{Data: &role}
	}()

	return output
}
