package repository

import "github.com/agungdwiprasetyo/agungdpcms/shared"

// User abstraction
type User interface {
	FindByUsername(username string) <-chan shared.Result
}

// Role abstraction
type Role interface {
	FindByID(id int) <-chan shared.Result
}
