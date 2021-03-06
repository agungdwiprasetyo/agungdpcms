package master

import (
	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	"github.com/agungdwiprasetyo/agungdpcms/src/master/delivery"
	"github.com/agungdwiprasetyo/agungdpcms/src/master/repository"
	"github.com/agungdwiprasetyo/agungdpcms/src/master/usecase"
)

// Module model
type Module struct {
	Handler    *delivery.GraphQLHandler
	Usecase    usecase.Master
	Repository *repository.Repository
}

// New master module constructor
func New(conf *config.Config, midd middleware.Middleware) *Module {
	repo := repository.NewRepository(conf.DB)
	uc := usecase.NewMasterUsecase(repo)
	handler := delivery.NewGraphQLHandler(uc, midd)

	return &Module{
		Handler:    handler,
		Usecase:    uc,
		Repository: repo,
	}
}
