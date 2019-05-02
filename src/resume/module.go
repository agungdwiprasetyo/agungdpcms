package resume

import (
	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/delivery"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/repository"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/usecase"
)

// Module model
type Module struct {
	Handler    *delivery.GraphQLHandler
	Usecase    usecase.Resume
	Repository *repository.Repository
}

// New resume module constructor
func New(conf *config.Config, midd middleware.Middleware) *Module {
	repo := repository.NewRepository(conf.DB)
	uc := usecase.NewResumeUsecase(repo)
	handler := delivery.NewGraphQLHandler(uc, midd)

	return &Module{
		Handler:    handler,
		Usecase:    uc,
		Repository: repo,
	}
}
