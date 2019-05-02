package delivery

import (
	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	"github.com/agungdwiprasetyo/agungdpcms/src/master/usecase"
)

// GraphQlHandler graphql handler for master module
type GraphQlHandler struct {
	uc   usecase.Master
	midd middleware.Middleware
}

// NewGraphQlHandler construct master delivery
func NewGraphQlHandler(uc usecase.Master, midd middleware.Middleware) *GraphQlHandler {
	return &GraphQlHandler{uc: uc, midd: midd}
}
