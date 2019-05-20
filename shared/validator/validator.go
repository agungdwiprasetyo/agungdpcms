package validator

import "github.com/agungdwiprasetyo/go-utils"

// Validator abstraction
type Validator interface {
	Validate(input interface{}) (multiError *utils.MultiError)
}
