package user

import (
	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	"github.com/agungdwiprasetyo/agungdpcms/shared/token"
	"github.com/agungdwiprasetyo/agungdpcms/src/user/delivery"
	"github.com/agungdwiprasetyo/agungdpcms/src/user/repository"
	"github.com/agungdwiprasetyo/agungdpcms/src/user/usecase"
)

// Module model
type Module struct {
	Handler    *delivery.GraphQLHandler
	Usecase    usecase.User
	Repository *repository.Repository
}

// New user module constructor
func New(conf *config.Config, midd middleware.Middleware, token token.Token) *Module {
	repo := repository.NewRepository(conf.DB)
	uc := usecase.NewUserUsecase(repo, token)
	handler := delivery.NewGraphQLHandler(uc, midd)

	return &Module{
		Handler:    handler,
		Usecase:    uc,
		Repository: repo,
	}
}
