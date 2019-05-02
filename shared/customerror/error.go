package customerror

import (
	"github.com/agungdwiprasetyo/go-utils"
)

// CustomError for custom query error
type CustomError interface {
	Error() string
	Extensions() map[string]interface{}
}

type customError struct {
	message    string
	multiError *utils.MultiError
}

// New custom error for validate request
func New(args ...interface{}) CustomError {
	if args == nil {
		return &customError{message: "error"}
	}

	ce := new(customError)
	for _, arg := range args {
		switch val := arg.(type) {
		case string:
			ce.message = val
		case *utils.MultiError:
			ce.multiError = val
		case utils.MultiError:
			ce.multiError = &val
		case error:
			if val != nil {
				ce.message = val.Error()
			}
		}
	}

	return ce
}

func (c *customError) Error() string {
	return c.message
}

func (c *customError) Extensions() (m map[string]interface{}) {
	if c.multiError != nil && c.multiError.HasError() {
		m = make(map[string]interface{})
		for k, v := range c.multiError.ToMap() {
			m[k] = v
		}
	}
	return
}
