package usecase

import "github.com/agungdwiprasetyo/agungdpcms/shared"

// User abstraction
type User interface {
	Login(username, password string) shared.Result
}
