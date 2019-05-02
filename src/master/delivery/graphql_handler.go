package delivery

import (
	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	"github.com/agungdwiprasetyo/agungdpcms/src/master/usecase"
)

// GraphQLHandler graphql handler for master module
type GraphQLHandler struct {
	uc   usecase.Master
	midd middleware.Middleware
}

// NewGraphQLHandler construct master delivery
func NewGraphQLHandler(uc usecase.Master, midd middleware.Middleware) *GraphQLHandler {
	return &GraphQLHandler{uc: uc, midd: midd}
}
